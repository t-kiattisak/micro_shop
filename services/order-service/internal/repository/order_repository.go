package repository

import (
	"order-service/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *domain.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}
