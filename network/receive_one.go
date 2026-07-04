package network

import (
	"bufio"
	"connect6/protocol"
	"encoding/json"
	"fmt"
	"net"
)

func ReceiveOneMessage(scanner *bufio.Scanner) (protocol.Message, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return protocol.Message{}, err
		}
		return protocol.Message{}, fmt.Errorf("connection closed before receiving a message")
	}

	var msg protocol.Message
	if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
		return protocol.Message{}, err
	}

	return msg, nil
}

func ReceiveMetadataResponse(conn net.Conn) (protocol.Message, error) {
	scanner := bufio.NewScanner(conn)

	for {
		msg, err := ReceiveOneMessage(scanner)
		if err != nil {
			return protocol.Message{}, err
		}

		switch msg.Type {
		case protocol.MessageTypeJoin:
			continue

		case protocol.MessageTypeMetadataResponse:
			return msg, nil
		}
	}
}