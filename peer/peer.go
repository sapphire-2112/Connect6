package peer

import (
	"net"
)
type Peer struct {
	ID        string
	Name      string
	Address   string
	Trusted   bool
	Connected bool
	Online    bool
	LastSeen  int64
	Conn      net.Conn
}