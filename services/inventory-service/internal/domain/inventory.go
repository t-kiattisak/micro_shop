package domain

import (
	"errors"
)

type Inventory struct {
	ID           uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Product      string  `gorm:"type:varchar(100);unique;not null" json:"product"`
	Quantity     int     `gorm:"not null" json:"quantity"`
	PricePerUnit float64 `gorm:"not null" json:"price_per_unit"`
}

func (i *Inventory) ReduceStock(qty int) error {
	if qty <= 0 {
		return errors.New("quantity to reduce must be greater than zero")
	}
	if i.Quantity < qty {
		return errors.New("insufficient stock")
	}
	i.Quantity -= qty
	return nil
}

func (i *Inventory) AddStock(qty int) error {
	if qty <= 0 {
		return errors.New("quantity to add must be greater than zero")
	}
	i.Quantity += qty
	return nil
}
