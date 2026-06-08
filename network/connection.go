package network

import (
	"crypto/tls"
	"net"
)

func Connect(address string) (net.Conn, error) {

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	return tls.Dial(
		"tcp6",
		address,
		config,
	)
}