package internal

import (
	"order-service/infrastructure"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/usecase"
)

func CreateOrderHandler() *handler.OrderHandler {
	db := infrastructure.ConnectDB()
	orderRepo := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)
	return handler.NewOrderHandler(orderUseCase)
}
