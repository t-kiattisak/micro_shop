package internal

import (
	"log"
	"payment-service/infrastructure"
	"payment-service/internal/domain"
	"payment-service/internal/grpcclient"
	"payment-service/internal/handler"
	"payment-service/internal/kafka"
	"payment-service/internal/repository"
	"payment-service/internal/usecase"
)

func CreatePaymentHandler() *handler.PaymentHandler {
	db := infrastructure.ConnectDB()

	log.Println("Running database migration...")
	err := db.AutoMigrate(&domain.Payment{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	paymentRepo := repository.NewPaymentRepository(db)
	orderClient := grpcclient.NewOrderClient()
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepo, orderClient)

	consumer := kafka.NewKafkaConsumer(paymentUseCase)
	go consumer.StartConsuming(3)

	return handler.NewPaymentHandler(paymentUseCase)
}
