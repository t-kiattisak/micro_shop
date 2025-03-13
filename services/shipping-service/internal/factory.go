package internal

import (
	"log"
	"shipping-service/infrastructure"
	"shipping-service/internal/domain"
	"shipping-service/internal/handler"
	"shipping-service/internal/kafka"
	"shipping-service/internal/repository"
	"shipping-service/internal/usecase"
)

func CreateShippingHandler() *handler.ShippingHandler {
	db := infrastructure.ConnectDB()

	log.Println("Running database migration...")
	err := db.AutoMigrate(&domain.Shipping{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	shippingRepo := repository.NewShippingRepository(db)
	shippingUseCase := usecase.NewShippingUseCase(shippingRepo)

	shippingConsumer := kafka.NewShippingConsumer(shippingUseCase)
	go shippingConsumer.StartConsuming(3)

	return handler.NewShippingHandler(shippingUseCase)
}
