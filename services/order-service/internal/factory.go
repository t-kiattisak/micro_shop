package internal

import (
	"order-service/infrastructure"
	"order-service/internal/grpcclient"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/usecase"
)

func CreateOrderHandler() *handler.OrderHandler {
	db := infrastructure.ConnectDB()
	orderRepo := repository.NewOrderRepository(db)
	inventoryClient := grpcclient.NewInventoryClient()
	orderUseCase := usecase.NewOrderUseCase(orderRepo, inventoryClient)
	return handler.NewOrderHandler(orderUseCase)
}
