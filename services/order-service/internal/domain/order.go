package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID     string  `gorm:"primaryKey" json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
