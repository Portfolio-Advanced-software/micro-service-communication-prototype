package main

import (
	messaging "github.com/Portfolio-Advanced-software/micro-service-communication-prototype/RabbitMQ/messaging"
)

func main() {
	messaging.ConsumeMessage("test-queue")
}
