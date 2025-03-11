package grpcclient

import (
	"context"
	"log"
	"time"

	"payment-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	client proto.OrderServiceClient
}

func NewOrderClient() *OrderClient {
	conn, err := grpc.NewClient(
		"localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to order-service: %v", err)
	}

	return &OrderClient{
		client: proto.NewOrderServiceClient(conn),
	}
}

func (c *OrderClient) CheckOrderExists(orderID uint) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.CheckOrderExists(ctx, &proto.CheckOrderRequest{OrderId: uint32(orderID)})
	if err != nil {
		return false, err
	}
	return resp.Exists, nil
}

func (c *OrderClient) UpdateOrderStatus(orderID uint, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.client.UpdateOrderStatus(ctx, &proto.UpdateOrderStatusRequest{
		OrderId: uint32(orderID),
		Status:  status,
	})
	if err != nil {
		log.Printf("❌ Failed to update order status: %v", err)
		return err
	}

	log.Printf("✅ Order %d marked as %s", orderID, status)
	return nil
}
