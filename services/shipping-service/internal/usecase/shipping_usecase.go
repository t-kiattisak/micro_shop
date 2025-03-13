package usecase

import (
	"log"
	"shipping-service/internal/domain"
	"shipping-service/internal/repository"
)

type ShippingUseCase struct {
	repo repository.ShippingRepository
}

func NewShippingUseCase(repo repository.ShippingRepository) *ShippingUseCase {
	return &ShippingUseCase{repo: repo}
}

func (uc *ShippingUseCase) CreateShipping(orderID uint, carrier string, trackingNumber string) error {
	shipping := &domain.Shipping{
		OrderID:        orderID,
		Carrier:        carrier,
		TrackingNumber: trackingNumber,
		Status:         "SHIPPING",
	}
	err := uc.repo.Create(shipping)
	if err != nil {
		log.Printf("Failed to create shipping record: %v", err)
	}
	return err
}

func (uc *ShippingUseCase) UpdateShippingStatus(orderID uint, status string) error {
	return uc.repo.UpdateStatus(orderID, status)
}
