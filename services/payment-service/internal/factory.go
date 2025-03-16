package internal

import (
	"log"
	"net"
	"payment-service/infrastructure"
	"payment-service/internal/domain"
	"payment-service/internal/grpc"
	"payment-service/internal/grpcclient"
	"payment-service/internal/handler"
	"payment-service/internal/kafka"
	"payment-service/internal/repository"
	"payment-service/internal/usecase"
	"payment-service/proto"

	grpcLib "google.golang.org/grpc"
)

func CreatePaymentHandler() *handler.PaymentHandler {
	db := infrastructure.ConnectDB()

	log.Println("Running database migration...")
	err := db.AutoMigrate(&domain.Payment{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	kafkaProducer := kafka.NewKafkaProducer("payment-events")

	paymentRepo := repository.NewPaymentRepository(db)
	orderClient := grpcclient.NewOrderClient()
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepo, orderClient, kafkaProducer)

	grpcServer := grpcLib.NewServer()
	grpcService := grpc.NewPaymentGRPCServer(paymentUseCase)

	proto.RegisterPaymentServiceServer(grpcServer, grpcService)
	go func() {
		log.Println("✅ gRPC Service is running on port 50053...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	consumerHandler := kafka.NewConsumerHandler(paymentUseCase)
	consumer := kafka.NewKafkaConsumer(consumerHandler)
	go consumer.StartConsuming(3)

	return handler.NewPaymentHandler(paymentUseCase)
}
