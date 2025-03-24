package grpcclient

import (
	"context"
	"log"
	"time"

	"order-service/proto"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	client  proto.InventoryServiceClient
	breaker *gobreaker.CircuitBreaker
}

func NewInventoryClient() *InventoryClient {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Failed to connect to inventory service: %v", err)
	}

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "PaymentServiceCB",
		MaxRequests: 5,
		Interval:    10 * time.Second,
		Timeout:     30 * time.Second,
	})

	return &InventoryClient{
		client:  proto.NewInventoryServiceClient(conn),
		breaker: cb,
	}
}

func (c *InventoryClient) CheckStock(product string, quantity int32) (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := c.breaker.Execute(func() (interface{}, error) {
		return c.client.CheckStock(ctx, &proto.CheckStockRequest{
			Product:  product,
			Quantity: quantity,
		})
	})

	if err != nil {
		log.Printf("❌ Circuit Breaker Open - Failed to check stock: %v", err)
		return false, "", err
	}

	resp := result.(*proto.CheckStockResponse)
	return resp.Available, resp.Message, nil
}

func (c *InventoryClient) ReduceStock(product string, quantity int32) (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := c.breaker.Execute(func() (interface{}, error) {
		return c.client.ReduceStock(ctx, &proto.ReduceStockRequest{
			Product:  product,
			Quantity: quantity,
		})
	})

	if err != nil {
		log.Printf("❌ Circuit Breaker Open - Failed to reduce stock: %v", err)
		return false, "", err
	}

	resp := result.(*proto.ReduceStockResponse)
	return resp.Success, resp.Message, nil
}

func (c *InventoryClient) GetPrice(product string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.GetPrice(ctx, &proto.GetPriceRequest{Product: product})
	if err != nil {
		return 0, err
	}
	price := float64(resp.Price)

	return price, nil
}
