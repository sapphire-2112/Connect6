package commands

import (
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"fmt"
	"time"
)

func Accept(
	manager *peer.Manager,
	targetID string,
	nodeID string,
) {

	p := manager.GetPeerByID(targetID)

	if p == nil {

		fmt.Println(
			"Peer not found",
		)

		return
	}

	msg := protocol.Message{
		Type: protocol.MessageTypeConnectionAccept,
		SenderID: nodeID,
		Timestamp: time.Now().Unix(),
	}

	network.SendMessage(
		p.Conn,
		msg,
	)

	fmt.Println(
		"Accepted request from",
		targetID,
	)
}