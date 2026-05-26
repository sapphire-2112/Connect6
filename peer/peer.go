package peer

import "net"

type Peer struct {
	ID      string
	Address string
	Conn    net.Conn
}