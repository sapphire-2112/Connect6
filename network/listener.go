package network

import (
	"crypto/tls"
	"net"
	"connect6/data"
)

func StartListener(address string) (net.Listener, error) {

	cert, err := tls.LoadX509KeyPair(
    data.CertPath(),
    data.KeyPath(),
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