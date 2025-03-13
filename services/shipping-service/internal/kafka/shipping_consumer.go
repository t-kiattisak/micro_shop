package kafka

import (
	"encoding/json"
	"log"
	"os"
	"shipping-service/internal/usecase"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ShippingConsumer struct {
	Consumer *kafka.Consumer
	usecase  *usecase.ShippingUseCase
}

func NewShippingConsumer(usecase *usecase.ShippingUseCase) *ShippingConsumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          "shipping-service",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("❌ Failed to create Kafka consumer: %v", err)
	}

	err = consumer.Subscribe("payment-events", nil)

	if err != nil {
		log.Fatalf("❌ Failed to subscribe to topic: %v", err)
	}

	return &ShippingConsumer{
		Consumer: consumer,
		usecase:  usecase,
	}
}

func (c *ShippingConsumer) StartConsuming(workerCount int) {
	log.Println("📥 Kafka Consumer for Shipping Service started...")
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			log.Printf("🚀 Worker %d started", workerID)

			for {
				event := c.Consumer.Poll(100)
				switch e := event.(type) {
				case *kafka.Message:
					log.Printf("📩 Worker %d received payment event: %s", workerID, string(e.Value))

					var msg struct {
						OrderID uint   `json:"order_id"`
						Status  string `json:"status"`
					}

					if err := json.Unmarshal(e.Value, &msg); err != nil {
						log.Printf("Worker %d failed to parse message: %v", workerID, err)
						PublishToDLQ("failed-payment-events", e.Value)
						continue
					}

					if msg.Status == "PAID" {
						log.Printf("🚚 Worker %d creating shipping for order %d", workerID, msg.OrderID)
						err := c.usecase.CreateShipping(msg.OrderID, "DHL", "TRACK12345")
						if err != nil {
							log.Printf("Worker %d failed to create shipping: %v", workerID, err)
						} else {
							log.Printf("Worker %d shipping for order %d created successfully!", workerID, msg.OrderID)
						}
					}

				case kafka.Error:
					if e.Code() != kafka.ErrTimedOut {
						log.Printf("Worker %d Kafka error: %v", workerID, e)
					} else {
						log.Printf("Worker %d waiting for new events...", workerID)
					}
				}
			}
		}(i + 1)
	}
	wg.Wait()
}
