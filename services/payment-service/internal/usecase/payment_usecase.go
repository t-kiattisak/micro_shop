package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"payment-service/internal/domain"
	"payment-service/internal/grpcclient"
	"payment-service/internal/repository"
)

type PaymentUseCase struct {
	repo           repository.PaymentRepository
	orderClient    *grpcclient.OrderClient
	eventPublisher PaymentEventPublisher
}

func NewPaymentUseCase(repo repository.PaymentRepository, orderClient *grpcclient.OrderClient, eventPublisher PaymentEventPublisher) *PaymentUseCase {
	return &PaymentUseCase{repo: repo,
		orderClient:    orderClient,
		eventPublisher: eventPublisher,
	}
}

func (uc *PaymentUseCase) CreatePayment(orderID uint, amount float64) (*domain.Payment, error) {
	exists, err := uc.orderClient.CheckOrderExists(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to verify order: %v", err)
	}
	if !exists {
		return nil, errors.New("order not found")
	}

	payment := &domain.Payment{
		OrderID: orderID,
		Amount:  amount,
		Status:  "PENDING",
	}
	err = uc.repo.Create(payment)
	return payment, err
}

func (uc *PaymentUseCase) GetPayment(id uint) (*domain.Payment, error) {
	return uc.repo.FindByID(id)
}

func (uc *PaymentUseCase) UpdatePaymentStatus(orderID uint, status string) error {
	payment, err := uc.repo.FindByOrderID(orderID)
	if err != nil {
		return err
	}

	payment.Status = status
	if err := uc.repo.Update(payment); err != nil {
		return err
	}

	if status == "PAID" {
		go uc.notifyOrderService(orderID)
	}

	return nil
}

func (uc *PaymentUseCase) notifyOrderService(orderID uint) {
	err := uc.orderClient.UpdateOrderStatus(orderID, "PAID")

	if err != nil {
		log.Printf("❌ Failed to notify order service: %v", err)
		return
	}
	log.Printf("✅ Order %d marked as PAID", orderID)
}

func (uc *PaymentUseCase) ProcessPayment(orderID uint, amount float64) error {
	exists, err := uc.orderClient.CheckOrderExists(orderID)
	if err != nil {
		return fmt.Errorf("failed to verify order: %v", err)
	}
	if !exists {
		return errors.New("order not found")
	}

	payment, err := uc.repo.FindByOrderID(orderID)
	if err != nil {
		payment = &domain.Payment{
			OrderID: orderID,
			Amount:  amount,
			Status:  "PROCESSING",
		}
		err = uc.repo.Create(payment)
		if err != nil {
			return fmt.Errorf("failed to create payment record: %v", err)
		}
	} else {
		payment.Status = "PROCESSING"
		if err := uc.repo.Update(payment); err != nil {
			return fmt.Errorf("failed to update payment status: %v", err)
		}

	}

	go uc.notifyOrderService(orderID)
	log.Printf("Payment for order %d completed successfully!", orderID)

	event := map[string]interface{}{
		"order_id": orderID,
		"amount":   amount,
		"status":   "PAID",
	}
	eventBytes, _ := json.Marshal(event)

	err = uc.eventPublisher.PublishMessage("payment-events", string(eventBytes))
	if err != nil {
		log.Printf("Failed to send event to Kafka: %v", err)
	}

	return nil
}
