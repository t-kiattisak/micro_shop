package usecase

import (
	"errors"
	"fmt"
	"order-service/internal/domain"
	"order-service/internal/grpcclient"
	"order-service/internal/repository"
)

type OrderUseCase struct {
	repo            repository.OrderRepository
	inventoryClient *grpcclient.InventoryClient
}

func NewOrderUseCase(repo repository.OrderRepository, inventoryClient *grpcclient.InventoryClient) *OrderUseCase {
	return &OrderUseCase{
		repo:            repo,
		inventoryClient: inventoryClient,
	}
}

func (u *OrderUseCase) CreateOrder(order *domain.Order) error {
	available, message, err := u.inventoryClient.CheckStock(order.Name, int32(order.Amount))
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

	success, msg, err := u.inventoryClient.ReduceStock(order.Name, int32(order.Amount))
	if err != nil || !success {
		return fmt.Errorf("order created, but stock reduction failed: %s", msg)
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
