package mobile

import (
	"connect6/commands"
	"connect6/data"
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"fmt"
	"log"
	"net"
	"time"
)

func (e *Engine) acceptLoop() {
	for {
		conn, err := e.listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		go e.handleConnection(conn)
	}
}

func (e *Engine) handleConnection(conn net.Conn) {

	p := &peer.Peer{
		ID:        "",
		Address:   conn.RemoteAddr().String(),
		Connected: true,
		Conn:      conn,
		Online:    true,
		LastSeen:  time.Now().Unix(),
	}

	join := protocol.Message{
		Type:     protocol.MessageTypeJoin,
		SenderID: e.identity.ID,

		PeerInfo: &protocol.PeerInfo{
			ID:        e.identity.ID,
			Name:      e.identity.Name,
			Address:   e.identity.Address,
			TrustedBy: e.identity.TrustedBy,
		},
		Timestamp: time.Now().Unix(),
	}

	network.SendMessage(conn, join)

	go network.ReceiveMessages(conn, func(msg protocol.Message) {

		if msg.Type == protocol.MessageTypeJoin {

			if p.ID != "" {
				return
			}

			if msg.PeerInfo != nil {

				p.ID = msg.PeerInfo.ID
				p.Name = msg.PeerInfo.Name
				p.Address = msg.PeerInfo.Address

			}
			e.manager.AddPeer(p)

			fmt.Println(
				"Identity Name and Address learned:",
				p.ID,
				p.Name,
				p.Address,
			)
			p.Online = true

			host, _, err := net.SplitHostPort(p.Address)
			if err != nil {
				fmt.Println("Address parse failed:", err)
				return
			}

			realAddress := net.JoinHostPort(host, "8080")

			storedPeer := data.StoredPeer{
				ID:           p.ID,
				Name:         p.Name,
				Address:      realAddress,
				Trusted:      false,
				TrustedSince: 0,
				TrustedBy:    msg.PeerInfo.TrustedBy,
				LastSeen:     p.LastSeen,
			}
			peers, _ := data.LoadPeers()
			exists := false

			fmt.Println(
				"Trusted By:",
				msg.PeerInfo.TrustedBy,
			)

			for _, peer := range peers {

				if peer.ID == storedPeer.ID {

					exists = true
					break
				}
			}
			if !exists {

				peers = append(peers, storedPeer)

				err = data.SavePeers(peers)
				if err != nil {
					fmt.Println("SavePeers failed:", err)
				}
			}

			return
		}
		if msg.Type == protocol.MessageTypePendingSyncRequest {

			pendingMessages, err := data.LoadPendingMessages()
			if err != nil {
				fmt.Println("Failed to load pending messages:", err)
			}

			for _, pending := range pendingMessages {

				if pending.ReceiverID != msg.SenderID {
					continue
				}

				chat := protocol.Message{
					Type:      protocol.MessageTypeChat,
					SenderID:  pending.SenderID,
					Payload:   pending.Content,
					Timestamp: pending.Timestamp,
				}

				err := network.SendMessage(conn, chat)
				if err != nil {
					fmt.Println("Failed to deliver pending message:", err)
					continue
				}

				err = data.RemovePendingMessage(pending.ID)
				if err != nil {
					fmt.Println("Failed to remove pending message:", err)
				}
			}
		}

		if msg.Type == protocol.MessageTypeReputationUpdate {

			peers, _ := data.LoadPeers()

			for i := range peers {

				if peers[i].ID == msg.PeerInfo.ID {

					peers[i].TrustedBy = msg.PeerInfo.TrustedBy
					break
				}
			}

			data.SavePeers(peers)

			return
		}

		if msg.Type == protocol.MessageTypePeerListRequest {

			storedPeers, err := data.LoadPeers()
			if err != nil {
				return
			}

			var peerInfos []protocol.PeerInfo

			for _, peer := range storedPeers {

				peerInfos = append(
					peerInfos,
					protocol.PeerInfo{
						ID:        peer.ID,
						Name:      peer.Name,
						Address:   peer.Address,
						TrustedBy: peer.TrustedBy,
					},
				)
			}

			response := protocol.Message{
				Type:      protocol.MessageTypePeerListResponse,
				SenderID:  e.identity.ID,
				PeerInfos: peerInfos,
				Timestamp: time.Now().Unix(),
			}

			network.SendMessage(
				conn,
				response,
			)

			return
		}

		if msg.Type == protocol.MessageTypePeerListResponse {

			fmt.Printf("\nPeers known by %s\n", msg.SenderID)

			for _, peer := range msg.PeerInfos {

				e.introducedPeers[peer.ID] = peer

				fmt.Printf("%s (%s)\n", peer.Name, peer.ID)

				info, err := commands.Metadata(peer, e.identity.ID)
				if err != nil {
					fmt.Println("Metadata unavailable:", err)
					continue
				}
				fmt.Printf("Known Peers   : %d\n", info.TotalContacts)
				fmt.Printf("Trust Ratio   : %.2f%%\n", info.TrustRatio)

				fmt.Printf("Trusted By: %d\n", info.TrustedBy)
				fmt.Printf("Confidence Score : %.2f\n", info.ConfidenceScore)
				fmt.Printf("Confidence       : %s\n", info.Confidence)

				fmt.Println("Mutual Trusted:")
				if len(info.MutualTrusted) == 0 {
					fmt.Println("  None")
				} else {
					for _, name := range info.MutualTrusted {
						fmt.Printf("  - %s\n", name)
					}
				}
			}
		}
		if msg.Type == protocol.MessageTypeConnectionRequest {

			e.pendingRequests[msg.PeerInfo.ID] = *msg.PeerInfo
			fmt.Println(
				"Pending count:",
				len(e.pendingRequests),
			)

			trustRatio := commands.CalculateTrustRatio(
				msg.PeerInfo.TrustedBy,
				msg.PeerInfo.TotalContacts,
			)

			confidenceScore := commands.CalculateConfidenceScore(
				msg.PeerInfo.TrustedBy,
				msg.PeerInfo.TotalContacts,
			)

			confidence := commands.GetConfidence(
				confidenceScore,
			)

			fmt.Printf(
				"\nConnection request received\n\n"+
					"Name: %s\n"+
					"ID: %s\n"+
					"Trusted By: %d\n"+
					"Known Peers: %d\n"+
					"Trust Ratio: %.2f%%\n"+
					"Confidence Score: %.2f\n"+
					"Confidence: %s\n",

				msg.PeerInfo.Name,
				msg.PeerInfo.ID,
				msg.PeerInfo.TrustedBy,
				msg.PeerInfo.TotalContacts,
				trustRatio,
				confidenceScore,
				confidence,
			)

			return
		}

		if msg.Type == protocol.MessageTypeConnectionAccept {

			if msg.PeerInfo == nil {
				return
			}

			fmt.Printf(
				"\n%s accepted your request\n",
				msg.PeerInfo.Name,
			)

			storedPeer := data.StoredPeer{
				ID:           msg.PeerInfo.ID,
				Name:         msg.PeerInfo.Name,
				Address:      msg.PeerInfo.Address,
				Trusted:      false,
				TrustedSince: 0,
				LastSeen:     time.Now().Unix(),
			}
			peers, _ := data.LoadPeers()
			exists := false

			for _, peer := range peers {

				if peer.ID == storedPeer.ID {

					exists = true
					break
				}
			}

			if !exists {

				peers = append(
					peers,
					storedPeer,
				)

				err := data.SavePeers(peers)

				if err != nil {
					fmt.Println(
						"SavePeers failed:",
						err,
					)
				}
			}
			fmt.Print("> ")

			return
		}

		if msg.Type == protocol.MessageTypeTrust {

			peers, _ := data.LoadPeers()
			identity, _ := data.LoadIdentity()

			for i := range peers {

				if peers[i].ID == msg.SenderID {

					peers[i].TrustsMe = true
					exists := false

					for _, id := range identity.TrustedByPeers {

						if id == msg.SenderID {

							exists = true
							break
						}
					}

					if !exists {

						identity.TrustedByPeers = append(
							identity.TrustedByPeers,
							msg.SenderID,
						)
					}

					break
				}
			}

			data.SavePeers(peers)

			identity.TrustedBy =
				data.GetTrustedByCount()

			data.SaveIdentity(identity)

			reputationUpdate := protocol.Message{
				Type:     protocol.MessageTypeReputationUpdate,
				SenderID: identity.ID,
				PeerInfo: &protocol.PeerInfo{
					ID:        identity.ID,
					TrustedBy: identity.TrustedBy,
				},
				Timestamp: time.Now().Unix(),
			}
			for _, peer := range e.manager.GetPeers() {

				if peer.Conn != nil {

					err := network.SendMessage(
						peer.Conn,
						reputationUpdate,
					)

					if err != nil {
						fmt.Println(
							"Failed to send reputation update:",
							err,
						)
					}
				}
			}

			fmt.Println(
				"Trust received from:",
				msg.SenderID,
			)
			fmt.Println(
				"My Trusted By:",
				identity.TrustedBy,
			)

			return
		}

		if msg.Type == protocol.MessageTypeMetadataRequest {

			identity, err := data.LoadIdentity()
			if err != nil {
				return
			}
			storedPeers, _ := data.LoadPeers()

			totalContacts := len(storedPeers)
			response := protocol.Message{
				Type:     protocol.MessageTypeMetadataResponse,
				SenderID: identity.ID,
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
			network.SendMessage(conn, response)

		}

		if msg.Type == protocol.MessageTypeHeartbeat {

			peers, _ := data.LoadPeers()

			for i := range peers {

				if peers[i].ID == msg.SenderID {

					peers[i].LastSeen = time.Now().Unix()

					break
				}
			}

			data.SavePeers(peers)

			return
		}

		chatMsg := data.ChatMessage{
			SenderID: msg.SenderID,
			Text:     msg.Payload,
			Time:     msg.Timestamp,
			Outgoing: false,
		}

		err := data.SaveMessage(
			msg.SenderID,
			chatMsg,
		)

		if err != nil {
			fmt.Println("Save failed:", err)
		}

		fmt.Printf(
			"\n[%s] %s\n",
			msg.SenderID,
			msg.Payload,
		)

		fmt.Print("> ")
	})

}
