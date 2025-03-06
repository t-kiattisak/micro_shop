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
