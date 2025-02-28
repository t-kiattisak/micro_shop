package domain

import (
	"gorm.io/gorm"
)

type Inventory struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Product  string `gorm:"type:varchar(100);unique;not null" json:"product"`
	Quantity int    `gorm:"not null" json:"quantity"`
}

func (i *Inventory) ReduceStock(qty int) error {
	if i.Quantity < qty {
		return gorm.ErrRecordNotFound
	}
	i.Quantity -= qty
	return nil
}

func (i *Inventory) AddStock(qty int) {
	i.Quantity += qty
}
