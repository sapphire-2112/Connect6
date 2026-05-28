package commands

import (
	"connect6/peer"
	"fmt"
)

func Disconnect(
	manager *peer.Manager,
	peerID string,
) {

	p := manager.GetPeerByID(peerID)

	if p == nil {
		fmt.Println("Peer not found")
		return
	}

	if !p.Connected {
		fmt.Println("Peer already disconnected")
		return
	}

	p.Conn.Close()

	p.Conn = nil
	p.Connected = false

	fmt.Println("Disconnected from", p.ID)
}