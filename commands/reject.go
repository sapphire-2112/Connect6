package commands

import (
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"fmt"
	"time"
)

func Reject(
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
		Type: protocol.MessageTypeConnectionReject,
		SenderID: nodeID,
		Timestamp: time.Now().Unix(),
	}

	network.SendMessage(
		p.Conn,
		msg,
	)

	fmt.Println(
		"Rejected request from",
		targetID,
	)
}