package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"order-service/internal/domain"
	"order-service/internal/grpcclient"
	"order-service/internal/kafka"
	"order-service/internal/repository"
)

type OrderUseCase struct {
	repo            repository.OrderRepository
	inventoryClient *grpcclient.InventoryClient
	kafkaProducer   *kafka.KafkaProducer
}

func NewOrderUseCase(repo repository.OrderRepository, inventoryClient *grpcclient.InventoryClient, kafkaProducer *kafka.KafkaProducer) *OrderUseCase {
	return &OrderUseCase{
		repo:            repo,
		inventoryClient: inventoryClient,
		kafkaProducer:   kafkaProducer,
	}
}

func (u *OrderUseCase) CreateOrder(order *domain.Order) error {

	pricePerUnit, err := u.inventoryClient.GetPrice(order.Name)
	if err != nil {
		return fmt.Errorf("failed to get product price: %v", err)
	}
	quantity := int(order.Amount / pricePerUnit)
	if order.Amount != float64(quantity)*pricePerUnit {
		return fmt.Errorf("invalid amount: %.2f is not a multiple of %.2f", order.Amount, pricePerUnit)
	}

	available, message, err := u.inventoryClient.CheckStock(order.Name, int32(quantity))
	if err != nil {
		return err
	}
	if !available {
		return fmt.Errorf("cannot create order: %s", message)
	}

	order.Status = "PENDING"
	err = u.repo.Create(order)
	if err != nil {
		return err
	}

	success, msg, err := u.inventoryClient.ReduceStock(order.Name, int32(quantity))
	if err != nil || !success {
		if _, delErr := u.repo.DeleteOrderById(order.ID); delErr != nil {
			return fmt.Errorf("order created, but stock reduction failed: %s, and failed to delete order: %v", msg, delErr)
		}
		return fmt.Errorf("order created, but stock reduction failed: %s", msg)
	}
	event := map[string]interface{}{
		"order_id": order.ID,
		"amount":   order.Amount,
		"status":   "PENDING",
	}
	eventBytes, _ := json.Marshal(event)

	err = u.kafkaProducer.PublishMessage(eventBytes)
	if err != nil {
		log.Printf("Failed to send event to Kafka: %v", err)
	}
	return nil
}

func (u *OrderUseCase) GetOrders() ([]domain.Order, error) {
	return u.repo.GetOrders()
}

func (u *OrderUseCase) GetOrderByID(id uint) (*domain.Order, error) {
	return u.repo.GetOrderByID(id)
}

func (u *OrderUseCase) DeleteOrderByID(id uint) (*domain.Order, error) {
	return u.repo.DeleteOrderById(id)
}

func (u *OrderUseCase) UpdateOrderStatus(id uint, newStatus string) error {
	validStatuses := map[string]bool{
		domain.OrderStatusPending:   true,
		domain.OrderStatusPaid:      true,
		domain.OrderStatusShipped:   true,
		domain.OrderStatusDelivered: true,
		domain.OrderStatusCancelled: true,
	}

	if !validStatuses[newStatus] {
		return errors.New("invalid status value")
	}

	order, err := u.repo.GetOrderByID(id)
	if err != nil {
		return err
	}

	order.Status = newStatus
	return u.repo.Update(order)
}

func (uc *OrderUseCase) CheckOrderExists(orderID uint) (bool, error) {
	order, err := uc.repo.GetOrderByID(orderID)
	if err != nil {
		return false, err
	}
	return order != nil, nil
}
