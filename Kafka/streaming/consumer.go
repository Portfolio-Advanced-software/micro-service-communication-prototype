package streaming

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// Set up signal handling to gracefully stop the consumer
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Read a single message from the Kafka topic
	message, err := consumer.FetchMessage(context.Background())
	if err != nil {
		log.Fatal("Failed to fetch message:", err)
	}

	// Process the received message
	fmt.Printf("Received message: %s\n", string(message.Value))

	// Close the consumer and release resources
	consumer.Close()

	// Send a termination signal to exit immediately
	signals <- syscall.SIGTERM
}
