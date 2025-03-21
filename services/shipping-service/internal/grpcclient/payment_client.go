package grpcclient

import (
	"context"
	"fmt"
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
	maxRetries := 3
	var err error

	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		_, err = c.client.UpdatePaymentStatus(ctx, &proto.UpdatePaymentStatusRequest{
			OrderId: uint32(orderID),
			Status:  status,
		})

		cancel()

		if err != nil {
			log.Printf("❌ Failed to update payment status: %v", err)
			return nil
		}

		log.Printf("🔁 Retry %d - Failed to update payment status: %v", i+1, err)
		time.Sleep(time.Duration(i+1) * time.Second)

	}

	return fmt.Errorf("max retries reached to update payment status for order %d: %v", orderID, err)
}
