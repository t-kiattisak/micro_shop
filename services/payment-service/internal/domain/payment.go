package domain

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `json:"order_id"`
	Amount    float64        `json:"amount"`
	Status    string         `json:"status"` // PENDING, PAID, FAILED
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
