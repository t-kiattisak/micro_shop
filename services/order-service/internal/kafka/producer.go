package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewKafkaProducer(topic string) *KafkaProducer {
	broker := os.Getenv("KAFKA_BROKER")

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	return &KafkaProducer{
		Producer: producer,
		Topic:    topic,
	}
}

func (p *KafkaProducer) PublishMessage(message []byte) error {
	deliveryChan := make(chan kafka.Event, 1)

	err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Failed to deliver message: %v", m.TopicPartition.Error)
	} else {
		log.Printf("Published message to topic: %s", p.Topic)
	}

	close(deliveryChan)
	return nil
}
