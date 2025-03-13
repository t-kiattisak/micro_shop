package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func PublishToDLQ(topic string, message []byte) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
	})
	if err != nil {
		log.Printf("Failed to create Kafka producer for DLQ: %v", err)
		return
	}
	defer producer.Close()

	deliveryChan := make(chan kafka.Event)

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
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
