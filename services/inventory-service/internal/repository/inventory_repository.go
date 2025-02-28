package repository

import (
	"inventory-service/internal/domain"

	"gorm.io/gorm"
)

type InventoryRepository interface {
	GetByProduct(product string) (*domain.Inventory, error)
	Update(inventory *domain.Inventory) error
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetByProduct(product string) (*domain.Inventory, error) {
	var inventory domain.Inventory
	err := r.db.Where("product = ?", product).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *inventoryRepository) Update(inventory *domain.Inventory) error {
	return r.db.Save(inventory).Error
}
