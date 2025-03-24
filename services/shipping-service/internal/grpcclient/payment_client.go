package grpcclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"shipping-service/proto"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentClient struct {
	client  proto.PaymentServiceClient
	breaker *gobreaker.CircuitBreaker
}

func NewPaymentClient() *PaymentClient {
	conn, err := grpc.NewClient(
		"localhost:50053",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("❌ Failed to connect to payment-service: %v", err)
	}

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "PaymentServiceCB",
		MaxRequests: 5,
		Interval:    10 * time.Second,
		Timeout:     30 * time.Second,
	})

	return &PaymentClient{client: proto.NewPaymentServiceClient(conn), breaker: cb}
}

func (c *PaymentClient) UpdatePaymentStatus(orderID uint, status string) error {
	maxRetries := 3
	var err error

	for i := 0; i < maxRetries; i++ {
		_, err := c.breaker.Execute(func() (interface{}, error) {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			_, err := c.client.UpdatePaymentStatus(ctx, &proto.UpdatePaymentStatusRequest{
				OrderId: uint32(orderID),
				Status:  status,
			})
			return nil, err
		})

		if err == nil {
			log.Printf("✅ Successfully updated payment status for order %d", orderID)
			return nil
		}

		log.Printf("🔁 Retry %d - Failed to update payment status: %v", i+1, err)
		time.Sleep(time.Duration(i+1) * time.Second)

	}

	return fmt.Errorf("max retries reached to update payment status for order %d: %v", orderID, err)
}
