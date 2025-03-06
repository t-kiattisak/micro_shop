package repository

import (
	"payment-service/internal/domain"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *domain.Payment) error
	FindByID(id uint) (*domain.Payment, error)
	Update(payment *domain.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(payment *domain.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) FindByID(id uint) (*domain.Payment, error) {
	var payment domain.Payment
	err := r.db.First(&payment, id).Error
	return &payment, err
}

func (r *paymentRepository) Update(payment *domain.Payment) error {
	return r.db.Save(payment).Error
}
