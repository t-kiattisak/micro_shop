package usecase

import (
	"log"
	"shipping-service/internal/domain"
	"shipping-service/internal/grpcclient"
	"shipping-service/internal/repository"
)

type ShippingUseCase struct {
	repo          repository.ShippingRepository
	paymentClient *grpcclient.PaymentClient
}

func NewShippingUseCase(repo repository.ShippingRepository, paymentClient *grpcclient.PaymentClient) *ShippingUseCase {
	return &ShippingUseCase{repo: repo, paymentClient: paymentClient}
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

func (uc *ShippingUseCase) UpdatePaymentStatus(orderID uint, status string) error {
	return uc.paymentClient.UpdatePaymentStatus(orderID, status)
}
