package network

import (
	"net"
)

func StartListener(address string) (net.Listener, error) {

	return net.Listen("tcp6", address)
}