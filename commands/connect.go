package commands

import (
	"connect6/network"
	"net"
)

func Connect(address string) (net.Conn, error) {

	return network.Connect(address)
}