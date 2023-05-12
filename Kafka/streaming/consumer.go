package streaming

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consume() {
	// Set up Kafka consumer configuration
	consumerConfig := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"}, // Kafka broker addresses
		Topic:   "test-topic",               // Topic to consume messages from
		GroupID: "my-consumer-group",        // Consumer group ID
	}

	// Create a new Kafka consumer
	consumer := kafka.NewReader(consumerConfig)

	// Continuously read messages from the Kafka topic
	for {
		message, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Failed to read message:", err)
		}

		// Process the received message
		fmt.Printf("Received message: %s\n", string(message.Value))
	}
}
