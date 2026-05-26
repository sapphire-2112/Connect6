package network

import (
	"bufio"
	"connect6/protocol"
	"encoding/json"
	"net"
)

func ReceiveMessages(
	conn net.Conn,
	handler func(protocol.Message),
) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		var msg protocol.Message

		err := json.Unmarshal(scanner.Bytes(), &msg)
		if err != nil {
			continue
		}

		handler(msg)
	}
}