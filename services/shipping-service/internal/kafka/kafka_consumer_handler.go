package kafka

import (
	"encoding/json"
	"log"
	"math/rand"
	"shipping-service/internal/usecase"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerHandler struct {
	usecase *usecase.ShippingUseCase
}

func NewConsumerHandler(usecase *usecase.ShippingUseCase) *ConsumerHandler {
	return &ConsumerHandler{usecase: usecase}
}

func (h *ConsumerHandler) HandleMessage(msg *kafka.Message) {
	log.Printf("📩 Received payment event: %s", string(msg.Value))

	var event struct {
		OrderID uint   `json:"order_id"`
		Status  string `json:"status"`
	}

	if err := json.Unmarshal(msg.Value, &event); err != nil {
		log.Printf("❌ Failed to parse message: %v", err)
		return
	}

	if event.Status != "PAID" {
		return
	}

	const maxRetries = 3
	var err error

	for i := 0; i < maxRetries; i++ {
		log.Printf("🚚 Creating shipping for order %d", event.OrderID)
		carrier := "DHL"
		trackingNumber := generateTrackingNumber(carrier)
		err = h.usecase.CreateShipping(event.OrderID, carrier, trackingNumber)

		if err == nil {
			err = h.usecase.UpdatePaymentStatus(event.OrderID, "COMPLETED")
			if err != nil {
				log.Printf("Felid Update payment status order %d", event.OrderID)
			} else {
				log.Printf("Update payment status order %d successfully!", event.OrderID)
			}

			return
		}

		log.Printf("🔁 Retry %d - Failed to create shipping: %v", i+1, err)
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	PublishToDLQ("failed-payment-events", msg.Value)
}

func generateTrackingNumber(carrier string) string {
	return carrier + "-" + randString(10)
}

func randString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
