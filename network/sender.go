package network

import (
	"connect6/protocol"
	"encoding/json"
	"net"
)

func SendMessage(conn net.Conn, msg protocol.Message) error {

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	_, err = conn.Write(append(data, '\n'))

	return err
}