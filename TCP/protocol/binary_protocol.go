package protocol

import (
	"encoding/binary"
	"io"
)

// Message represents the binary message structure
type Message struct {
	ID      uint32
	Payload []byte
}

// ReadMessage reads a binary message from the given reader
func ReadMessage(reader io.Reader) (*Message, error) {
	header := make([]byte, 4)
	_, err := io.ReadFull(reader, header)
	if err != nil {
		return nil, err
	}

	messageLength := binary.BigEndian.Uint32(header)
	payload := make([]byte, messageLength)
	_, err = io.ReadFull(reader, payload)
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:      binary.BigEndian.Uint32(payload[:4]),
		Payload: payload[4:],
	}, nil
}

// WriteMessage writes a binary message to the given writer
func WriteMessage(writer io.Writer, message *Message) error {
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(len(message.Payload)+4))
	_, err := writer.Write(header)
	if err != nil {
		return err
	}

	payload := make([]byte, 4+len(message.Payload))
	binary.BigEndian.PutUint32(payload, message.ID)
	copy(payload[4:], message.Payload)
	_, err = writer.Write(payload)
	if err != nil {
		return err
	}

	return nil
}
