package kafka

import (
	"context"
	"encoding/json"
	"log"
	"payment-service/internal/usecase"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader  *kafka.Reader
	usecase *usecase.PaymentUseCase
}

func NewKafkaConsumer(brokerAddress, topic string, usecase *usecase.PaymentUseCase) *KafkaConsumer {
	return &KafkaConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{brokerAddress},
			Topic:    topic,
			GroupID:  "payment-service",
			MaxBytes: 10e6,
		}),
		usecase: usecase,
	}
}

func (c *KafkaConsumer) StartCumming() {
	for {
		msg, err := c.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			continue
		}
		var event map[string]interface{}
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			continue
		}

		orderID, ok := event["order_id"].(float64)
		if !ok {
			log.Printf("❌ Invalid order_id format")
			continue
		}

		amount, ok := event["amount"].(float64)
		if !ok {
			log.Printf("❌ Invalid amount format, using default = 0")
			amount = 0
		}

		log.Printf("📩 Received order event: OrderID=%d, Amount=%.2f", uint(orderID), amount)

		c.usecase.ProcessPayment(uint(orderID), amount)
	}
}
