package usecase

import (
	"errors"
	"order-service/internal/domain"
	"order-service/internal/repository"
)

type OrderUseCase struct {
	repo repository.OrderRepository
}

func NewOrderUseCase(repo repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{repo: repo}
}

func (u *OrderUseCase) CreateOrder(order *domain.Order) error {
	order.Status = "PENDING"
	return u.repo.Create(order)
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
