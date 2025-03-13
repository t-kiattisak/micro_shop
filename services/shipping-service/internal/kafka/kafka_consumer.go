package kafka

import (
	"log"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ShippingConsumer struct {
	Consumer *kafka.Consumer
	handler  *ConsumerHandler
}

func NewShippingConsumer(handler *ConsumerHandler) *ShippingConsumer {
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
		handler:  handler,
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
					c.handler.HandleMessage(e)

				case kafka.Error:
					if e.Code() != kafka.ErrTimedOut {
						log.Printf("❌ Worker %d Kafka error: %v", workerID, e)
					} else {
						log.Printf("⌛ Worker %d waiting for new events...", workerID)
					}
				}
			}
		}(i + 1)
	}
	wg.Wait()
}
