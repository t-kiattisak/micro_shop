package repository

import (
	"shipping-service/internal/domain"

	"gorm.io/gorm"
)

type ShippingRepository interface {
	Create(shipping *domain.Shipping) error
	GetByOrderID(orderID uint) (*domain.Shipping, error)
	UpdateStatus(orderID uint, status string) error
}

type shippingRepository struct {
	db *gorm.DB
}

func NewShippingRepository(db *gorm.DB) ShippingRepository {
	return &shippingRepository{db: db}
}

func (r *shippingRepository) Create(shipping *domain.Shipping) error {
	return r.db.Create(shipping).Error
}

func (r *shippingRepository) GetByOrderID(orderID uint) (*domain.Shipping, error) {
	var shipping domain.Shipping
	err := r.db.Where("order_id = ?", orderID).First(&shipping).Error
	return &shipping, err
}

func (r *shippingRepository) UpdateStatus(orderID uint, status string) error {
	return r.db.Model(&domain.Shipping{}).Where("order_id = ?", orderID).
		Update("status", status).Error
}
