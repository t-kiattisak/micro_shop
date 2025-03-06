package usecase

import (
	"errors"
	"fmt"
	"payment-service/internal/domain"
	"payment-service/internal/grpcclient"
	"payment-service/internal/repository"
)

type PaymentUseCase struct {
	repo        repository.PaymentRepository
	orderClient *grpcclient.OrderClient
}

func NewPaymentUseCase(repo repository.PaymentRepository, orderClient *grpcclient.OrderClient) *PaymentUseCase {
	return &PaymentUseCase{repo: repo,
		orderClient: orderClient,
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

func (uc *PaymentUseCase) UpdatePaymentStatus(id uint, status string) error {
	payment, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	if payment.Status != "PENDING" {
		return errors.New("cannot update non-pending payment")
	}
	payment.Status = status
	return uc.repo.Update(payment)
}
