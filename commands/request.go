package commands

import (
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
){

	peerInfo, exists :=
	introducedPeers[targetID]

		if !exists {

			fmt.Println(
				"Peer not known",
			)

			return
		}
		conn, err := Connect(
	peerInfo.Address,
			)

			if err != nil {

				fmt.Println(
					"Connection failed",
				)

				return
			}


			msg := protocol.Message{
				Type: protocol.MessageTypeConnectionRequest,

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
					"Request failed",
				)

				return
			}
			conn.Close()

			fmt.Printf(
				"Request sent to %s (%s)\n",
				peerInfo.Name,
				peerInfo.ID,
			)

	
	
}