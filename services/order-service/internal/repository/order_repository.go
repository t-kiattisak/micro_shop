package repository

import (
	"order-service/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	GetOrders() ([]domain.Order, error)
	GetOrderByID(id uint) (*domain.Order, error)
	DeleteOrderById(id uint) (*domain.Order, error)
	Update(order *domain.Order) error
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

func (r *orderRepository) GetOrders() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetOrderByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) DeleteOrderById(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.Delete(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}
