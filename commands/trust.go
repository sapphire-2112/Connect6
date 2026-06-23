package commands

import (
	"connect6/data"
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"fmt"
	"time"
)

func Trust(targetID string, nodeID string, manager *peer.Manager) {

	peers, err := data.LoadPeers()

	if err != nil {
		fmt.Println("LoadPeers failed")
		return
	}

	found := false

	for i := range peers {

		if peers[i].ID == targetID {

			peers[i].Trusted = true
			peers[i].TrustedSince = time.Now().Unix()

			found = true
			break
		}
	}

	if !found {

		fmt.Println("Peer not found")
		return
	}

	err = data.SavePeers(peers)

	if err != nil {

		fmt.Println("SavePeers failed")
		return
	}
		p := manager.GetPeerByID(
		targetID,
	)

	if p != nil {
		msg := protocol.Message{
    Type: protocol.MessageTypeTrust,
    SenderID: nodeID,
    Timestamp: time.Now().Unix(),
}
network.SendMessage(p.Conn, msg)
	}

	fmt.Println("Trusted:", targetID)
}