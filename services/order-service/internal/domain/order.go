package domain

import (
	"time"

	"gorm.io/gorm"
)

const (
	OrderStatusPending   = "PENDING"
	OrderStatusPaid      = "PAID"
	OrderStatusShipped   = "SHIPPED"
	OrderStatusDelivered = "DELIVERED"
	OrderStatusCancelled = "CANCELLED"
)

type Order struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `json:"name"`
	Amount    float64        `json:"amount"`
	Status    string         `gorm:"type:varchar(20)" json:"status"`
	Product   string         `gorm:"type:varchar(100);not null" json:"product"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
