package grpcclient

import (
	"context"
	"log"
	"time"

	"shipping-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentClient struct {
	client proto.PaymentServiceClient
}

func NewPaymentClient() *PaymentClient {
	conn, err := grpc.NewClient(
		"localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("❌ Failed to connect to payment-service: %v", err)
	}
	return &PaymentClient{client: proto.NewPaymentServiceClient(conn)}
}

func (c *PaymentClient) UpdatePaymentStatus(orderID uint, status string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := c.client.UpdatePaymentStatus(ctx, &proto.UpdatePaymentStatusRequest{
		OrderId: uint32(orderID),
		Status:  status,
	})

	if err != nil {
		log.Printf("❌ Failed to update payment status: %v", err)
		return err
	}

	return nil
}
