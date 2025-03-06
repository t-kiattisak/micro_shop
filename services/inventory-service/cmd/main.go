package main

import (
	"log"
	"net"

	"inventory-service/infrastructure"
	"inventory-service/internal/domain"
	"inventory-service/internal/grpc"
	"inventory-service/internal/handler"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"inventory-service/proto"

	fiber "github.com/gofiber/fiber/v2"
	grpcLib "google.golang.org/grpc"
)

func main() {
	db := infrastructure.ConnectDB()

	repo := repository.NewInventoryRepository(db)
	usecase := usecase.NewInventoryUseCase(repo)

	log.Println("Running database migration...")
	err := db.AutoMigrate(&domain.Inventory{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("✅ Database migration completed successfully!")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpcLib.NewServer()
	grpcService := grpc.NewInventoryGRPCServer(usecase)

	proto.RegisterInventoryServiceServer(grpcServer, grpcService)

	go func() {
		log.Println("✅ gRPC Service is running on port 50051...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	app := fiber.New()
	handler := handler.NewInventoryHandler(usecase)

	app.Post("/inventory", handler.CreateInventory)
	app.Get("/inventory/check", handler.CheckStock)
	app.Post("/inventory/reduce", handler.ReduceStock)

	log.Println("✅ Inventory Service is running on port 8082...")
	log.Fatal(app.Listen(":8082"))
}
