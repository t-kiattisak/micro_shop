package kafka

import (
	"log"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	Consumer        *kafka.Consumer
	consumerHandler *ConsumerHandler
}

func NewKafkaConsumer(consumerHandler *ConsumerHandler) *KafkaConsumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          "payment-service",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	err = consumer.Subscribe("order-events", nil)
	if err != nil {
		log.Fatalf("❌ Failed to subscribe to topic: %v", err)
	}

	return &KafkaConsumer{
		Consumer:        consumer,
		consumerHandler: consumerHandler,
	}
}

func (c *KafkaConsumer) StartConsuming(workerCount int) {
	log.Println("📥 Kafka Consumer for Payment Service started...")
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			log.Printf("Worker %d started", workerID)

			for {
				event := c.Consumer.Poll(100)
				switch e := event.(type) {
				case *kafka.Message:
					log.Printf("Worker %d received order event: %s", workerID, string(e.Value))

					c.consumerHandler.ProcessPaymentMessage(e.Value)

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
