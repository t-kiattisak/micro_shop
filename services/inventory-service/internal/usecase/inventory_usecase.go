package usecase

import (
	"errors"
	"inventory-service/internal/domain"
	"inventory-service/internal/repository"
)

type InventoryUseCase struct {
	repo repository.InventoryRepository
}

func NewInventoryUseCase(repo repository.InventoryRepository) *InventoryUseCase {
	return &InventoryUseCase{repo: repo}
}

func (u *InventoryUseCase) CheckStock(product string, qty int) error {
	inventory, err := u.repo.GetByProduct(product)
	if err != nil {
		return err
	}

	if inventory.Quantity < qty {
		return errors.New("not enough stock")
	}
	return nil
}

func (u *InventoryUseCase) CreateInventory(product string, qty int) error {
	existing, err := u.repo.GetByProduct(product)
	if err == nil && existing != nil {
		return errors.New("product already exists in inventory")
	}

	newInventory := &domain.Inventory{
		Product:  product,
		Quantity: qty,
	}
	return u.repo.Create(newInventory)
}

func (u *InventoryUseCase) ReduceStock(product string, qty int) error {
	inventory, err := u.repo.GetByProduct(product)
	if err != nil {
		return err
	}

	if err := inventory.ReduceStock(qty); err != nil {
		return err
	}

	return u.repo.Update(inventory)
}
