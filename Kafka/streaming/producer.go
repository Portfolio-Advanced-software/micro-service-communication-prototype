package streaming

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce() {
	// Set up Kafka producer configuration
	producerConfig := kafka.WriterConfig{
		Brokers: []string{"localhost:9092"}, // Kafka broker addresses
		Topic:   "test-topic",               // Topic to produce messages to
	}

	// Create a new Kafka producer
	producer := kafka.NewWriter(producerConfig)

	// Start producing a stream of messages
	for i := 1; ; i++ {
		message := kafka.Message{
			Value: []byte("Message " + strconv.Itoa(i)),
		}

		// Write the message to the Kafka topic
		err := producer.WriteMessages(context.Background(), message)
		if err != nil {
			log.Fatal("Failed to write message:", err)
		}

		time.Sleep(1 * time.Second) // Delay between producing messages
	}
}
