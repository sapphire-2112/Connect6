package network

import (
	"net"
)

func Connect(address string) (net.Conn, error) {

	return net.Dial("tcp6", address)
}