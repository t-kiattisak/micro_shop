package kafka

import (
	"log"
	"os"
	"time"

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

func (p *KafkaProducer) PublishMessage(key, message string) error {
	deliveryChan := make(chan kafka.Event, 1)

	err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(message),
		Timestamp:      time.Now(),
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
		log.Printf("Published message to topic: %s, key: %s", p.Topic, key)
	}

	close(deliveryChan)
	return nil
}

func PublishToDLQ(topic string, message []byte) {
	broker := os.Getenv("KAFKA_BROKER")

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		log.Printf("❌ Failed to create Kafka producer for DLQ: %v", err)
		return
	}
	defer producer.Close()

	deliveryChan := make(chan kafka.Event, 1)

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("DLQ"),
		Value:          message,
		Timestamp:      time.Now(),
	}, deliveryChan)

	if err != nil {
		log.Printf("Failed to send message to DLQ: %v", err)
		return
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Failed to deliver message to DLQ: %v", m.TopicPartition.Error)
	} else {
		log.Printf("Sent message to DLQ: %s", topic)
	}

	close(deliveryChan)
}
