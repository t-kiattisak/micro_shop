package usecase

import (
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
