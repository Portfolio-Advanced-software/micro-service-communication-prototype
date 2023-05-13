package main

import (
	"fmt"
	"net"

	"github.com/Portfolio-Advanced-software/micro-service-communication-prototype/TCP/protocol"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	message := &protocol.Message{
		ID:      42,
		Payload: []byte("Hello, server!"),
	}

	err = protocol.WriteMessage(conn, message)
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return
	}

	fmt.Println("Message sent successfully!")
}
