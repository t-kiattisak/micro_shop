package grpcclient

import (
	"context"
	"log"
	"time"

	"order-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InventoryClient struct {
	client proto.InventoryServiceClient
}

func NewInventoryClient() *InventoryClient {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("Failed to connect to inventory service: %v", err)
	}
	return &InventoryClient{
		client: proto.NewInventoryServiceClient(conn),
	}
}

func (c *InventoryClient) CheckStock(product string, quantity int32) (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.CheckStock(ctx, &proto.CheckStockRequest{
		Product:  product,
		Quantity: quantity,
	})
	if err != nil {
		return false, "", err
	}
	return resp.Available, resp.Message, nil
}

func (c *InventoryClient) ReduceStock(product string, quantity int32) (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.ReduceStock(ctx, &proto.ReduceStockRequest{
		Product:  product,
		Quantity: quantity,
	})
	if err != nil {
		return false, "", err
	}
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
