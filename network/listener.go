package network

import (
	"crypto/tls"
	"net"
)

func StartListener(address string) (net.Listener, error) {

	cert, err := tls.LoadX509KeyPair(
		"data/cert.pem",
		"data/key.pem",
	)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return tls.Listen(
		"tcp6",
		address,
		config,
	)
}