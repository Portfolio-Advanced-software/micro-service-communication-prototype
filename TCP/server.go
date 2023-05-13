package main

import (
	"fmt"
	"net"

	"github.com/Portfolio-Advanced-software/micro-service-communication-prototype/TCP/protocol"
)

func handleConnection(conn net.Conn) {
	message, err := protocol.ReadMessage(conn)
	if err != nil {
		fmt.Println("Failed to read message:", err)
		return
	}

	fmt.Println("Received message ID:", message.ID)
	fmt.Println("Received message payload:", string(message.Payload))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started, waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
