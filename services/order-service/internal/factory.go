package internal

import (
	"log"
	"net"
	"order-service/infrastructure"
	"order-service/internal/grpc"
	"order-service/internal/grpcclient"
	"order-service/internal/handler"
	"order-service/internal/kafka"
	"order-service/internal/repository"
	"order-service/internal/usecase"

	"order-service/proto"

	grpcLib "google.golang.org/grpc"
)

func CreateOrderHandler() *handler.OrderHandler {
	db := infrastructure.ConnectDB()
	orderRepo := repository.NewOrderRepository(db)
	inventoryClient := grpcclient.NewInventoryClient()

	kafkaProducer := kafka.NewKafkaProducer("order-events")

	orderUseCase := usecase.NewOrderUseCase(orderRepo, inventoryClient, kafkaProducer)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpcLib.NewServer()
	grpcService := grpc.NewOrderGRPCServer(orderUseCase)

	proto.RegisterOrderServiceServer(grpcServer, grpcService)
	go func() {
		log.Println("✅ gRPC Service is running on port 50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	return handler.NewOrderHandler(orderUseCase)
}
