package commands

import (
	"connect6/network"
	"connect6/protocol"
	"connect6/data"
	"fmt"
	"time"
)

func Accept(
	targetID string,
	nodeID string,
	nodeName string,
	nodeAddress string,
	pendingRequests map[string]protocol.PeerInfo,
) {

	requestInfo, rexists := pendingRequests[targetID]

	if !rexists {

			fmt.Println(
				"Request not found",
			)

			return
		}

	storedPeer := data.StoredPeer{
	ID:       requestInfo.ID,
	Name:     requestInfo.Name,
	Address:  requestInfo.Address,
	Trusted:  false,
	LastSeen: time.Now().Unix(),
}

	peers, _ := data.LoadPeers()


	exists := false

	for _, peer := range peers {

		if peer.ID == storedPeer.ID {

			exists = true
			break
		}
	}
	if(!exists){
	peers = append(peers, storedPeer)

	err:=data.SavePeers(peers)
	 if err != nil {
        fmt.Println(
            "SavePeers failed:",
            err,
        )
    }

	}

	conn, err := Connect(
	requestInfo.Address,
)

if err != nil {

	fmt.Println(
		"Connection failed",
	)

	return
}
msg := protocol.Message{
	Type: protocol.MessageTypeConnectionAccept,

	SenderID: nodeID,

	PeerInfo: &protocol.PeerInfo{
		ID:      nodeID,
		Name:    nodeName,
		Address: nodeAddress,
	},

	Timestamp: time.Now().Unix(),
}
err = network.SendMessage(
	conn,
	msg,
)

if err != nil {

	conn.Close()

	fmt.Println(
		"Accept failed",
	)

	return
}
conn.Close()
delete(
	pendingRequests,
	targetID,
)

fmt.Println(
	"Accepted:",
	requestInfo.Name,
)
}