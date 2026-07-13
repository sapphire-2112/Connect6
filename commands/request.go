package commands

import (
	"connect6/data"
	"connect6/network"
	"connect6/protocol"
	"fmt"
	"time"
)

func Request(
	targetID string,
	nodeID string,
	nodeName string,
	nodeAddress string,
	introducedPeers map[string]protocol.PeerInfo,
) {

	peerInfo, exists := introducedPeers[targetID]

	if !exists {
		fmt.Println("Peer not known")
		return
	}

	conn, err := Connect(peerInfo.Address)
	if err != nil {
		fmt.Println("Connection failed")
		return
	}
	defer conn.Close()

	identity, err := data.LoadIdentity()
	if err != nil {
		fmt.Println("Failed to load identity")
		return
	}

	storedPeers, err := data.LoadPeers()
	if err != nil {
		fmt.Println("Failed to load peers")
		return
	}

	totalContacts := len(storedPeers)

	msg := protocol.Message{
		Type:     protocol.MessageTypeConnectionRequest,
		SenderID: nodeID,
		PeerInfo: &protocol.PeerInfo{
			ID:             identity.ID,
			Name:           identity.Name,
			Address:        identity.Address,
			TrustedBy:      identity.TrustedBy,
			TrustedByPeers: identity.TrustedByPeers,
			TotalContacts:  totalContacts,
		},
		Timestamp: time.Now().Unix(),
	}

	err = network.SendMessage(conn, msg)
	if err != nil {
		fmt.Println("Request failed")
		return
	}

	fmt.Printf(
		"Request sent to %s (%s)\n",
		peerInfo.Name,
		peerInfo.ID,
	)
}