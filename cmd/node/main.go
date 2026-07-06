package main

import (
	"bufio"
	"connect6/commands"
	"connect6/crypto"
	"connect6/data"
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var manager = peer.NewManager()
var currentTarget string
var nodeID string
var nodeAddress string
var nodeName string
var introducedPeers = make(map[string]protocol.PeerInfo)
var pendingRequests = make(map[string]protocol.PeerInfo)

func handleConnection(conn net.Conn) {

	p := &peer.Peer{
		ID:        "",
		Address:   conn.RemoteAddr().String(),
		Connected: true,
		Conn:      conn,
		Online:    true,
		LastSeen:  time.Now().Unix(),
	}

	//fmt.Println("Connected:", p.Address)
	identity, _ := data.LoadIdentity()

	join := protocol.Message{
		Type:     protocol.MessageTypeJoin,
		SenderID: nodeID,

		PeerInfo: &protocol.PeerInfo{
			ID:        nodeID,
			Name:      nodeName,
			Address:   nodeAddress,
			TrustedBy: identity.TrustedBy,
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
			manager.AddPeer(p)

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
				SenderID:  nodeID,
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

				introducedPeers[peer.ID] = peer

				fmt.Printf("%s (%s)\n", peer.Name, peer.ID)

				info, err := commands.Metadata(peer, nodeID)
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

			pendingRequests[msg.PeerInfo.ID] = *msg.PeerInfo
			fmt.Println(
				"Pending count:",
				len(pendingRequests),
			)

			fmt.Printf(
				"\nConnection request received\n\n"+
					"Name: %s\n"+
					"ID: %s\n",

				msg.PeerInfo.Name,
				msg.PeerInfo.ID,
			)

			fmt.Println(
				"\nUse:\n/accept " +
					msg.PeerInfo.ID +
					"\n/reject " +
					msg.PeerInfo.ID,
			)

			fmt.Print("> ")

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
			for _, peer := range manager.GetPeers() {

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

func main() {

	nameFlag := flag.String(
		"n",
		"",
		"node name",
	)

	addressFlag := flag.String(
		"a",
		"",
		"advertised address",
	)

	flag.Parse()

	identity, err := data.LoadIdentity()
	if err != nil {
		fmt.Println(err)
		return
	}
	if identity.Name == "" {

		if *nameFlag == "" {

			fmt.Println(
				"Node name required",
			)

			return
		}

		identity.Name = *nameFlag

		err = data.SaveIdentity(
			identity,
		)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if identity.Address == "" {

		if *addressFlag == "" {

			fmt.Println(
				"Advertised address required",
			)

			return
		}

		identity.Address = *addressFlag
	}

	err = data.SaveIdentity(identity)

	if err != nil {
		fmt.Println(err)
		return
	}

	nodeName = identity.Name
	nodeAddress = identity.Address

	nodeID = identity.ID
	fmt.Println("Node ID:", nodeID)
	fmt.Println("Node Name:", identity.Name)
	fmt.Println("Node Address:", identity.Address)

	storedPeers, err := data.LoadPeers()
	go func() {

		for {

			time.Sleep(15 * time.Second)

			peers := manager.GetPeers()

			for _, p := range peers {

				if p.Online {
					continue
				}

				conn, err := commands.Connect(p.Address)

				if err != nil {
					continue
				}
				p.Online = true

				fmt.Println("Reconnected:", p.ID)

				go handleConnection(conn)
			}
		}
	}()

	if err == nil {

		for _, sp := range storedPeers {

			p := &peer.Peer{
				ID:        sp.ID,
				Address:   sp.Address,
				LastSeen:  sp.LastSeen,
				Online:    false,
				Connected: false,
			}

			manager.AddPeer(p)
		}

		fmt.Println(
			"Loaded peers:",
			len(storedPeers),
		)
	}

	if _, err := os.Stat("data/cert.pem"); os.IsNotExist(err) {

		err := crypto.GenerateCertificate(
			"data/cert.pem",
			"data/key.pem",
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("TLS certificate generated")
	}

	listener, err := network.StartListener("[::]:8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	fmt.Println("Node listening on port 8080")

	go func() {

		for {

			conn, err := listener.Accept()
			if err != nil {
				continue
			}

			go handleConnection(conn)
		}
	}()
	//this is heartbeat func
	go func() {

		for {

			time.Sleep(10 * time.Second)

			peers, err := data.LoadPeers()

			if err != nil {
				continue
			}

			for _, peer := range peers {

				conn, err := commands.Connect(
					peer.Address,
				)

				if err != nil {
					continue
				}

				msg := protocol.Message{
					Type:      protocol.MessageTypeHeartbeat,
					SenderID:  nodeID,
					Timestamp: time.Now().Unix(),
				}

				network.SendMessage(
					conn,
					msg,
				)

				conn.Close()
			}
		}
	}()

	// go func() {

	// 	for {

	// 		time.Sleep(5 * time.Second)

	// 		peers := manager.GetPeers()

	// 		for _, p := range peers {

	// 			if time.Now().Unix()-p.LastSeen > 30 {

	// 				p.Online = false
	// 				  if p.Conn != nil {
	// 						p.Conn.Close()
	// 					}
	// 				}
	// 		}
	// 	}
	// }()

	input := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")

		if !input.Scan() {
			return
		}

		text := input.Text()
		if text == "" {
			continue
		}

		if strings.HasPrefix(
			text,
			"/request ",
		) {

			targetID := strings.TrimPrefix(
				text,
				"/request ",
			)

			commands.Request(
				targetID,
				nodeID,
				nodeName,
				nodeAddress,
				introducedPeers,
			)

			continue
		}

		if strings.HasPrefix(text, "/accept ") {

			targetID := strings.TrimPrefix(
				text,
				"/accept ",
			)

			commands.Accept(
				targetID,
				nodeID,
				nodeName,
				nodeAddress,
				pendingRequests,
			)

			continue
		}

		if strings.HasPrefix(text, "/reject ") {

			targetID := strings.TrimPrefix(
				text,
				"/reject ",
			)

			commands.Reject(
				manager,
				targetID,
				nodeID,
			)

			continue
		}

		if strings.HasPrefix(text, "/connect ") {

			address := strings.TrimPrefix(text, "/connect ")

			conn, err := commands.Connect(address)
			if err != nil {
				fmt.Println("Connection failed")
				continue
			}
			join := protocol.Message{
				Type:     protocol.MessageTypeJoin,
				SenderID: nodeID,
				PeerInfo: &protocol.PeerInfo{
					ID:      nodeID,
					Name:    identity.Name,
					Address: nodeAddress,
				},

				Timestamp: time.Now().Unix(),
			}

			network.SendMessage(conn, join)

			go handleConnection(conn)

			fmt.Println("Connected to", address)

			continue
		}
		///Added Peers Intro

		if strings.HasPrefix(
			text,
			"/peersof ",
		) {

			targetID := strings.TrimPrefix(
				text,
				"/peersof ",
			)

			commands.PeersOf(
				manager,
				targetID,
				nodeID,
			)

			continue
		}

		//disconnect
		if strings.HasPrefix(text, "/disconnect ") {

			targetID := strings.TrimPrefix(text, "/disconnect ")

			commands.Disconnect(manager, targetID)

			continue
		}

		// SHOW PEERS
		if text == "/peers" {

			storedPeers, err := data.LoadPeers()

			if err != nil {
				fmt.Println(
					"LoadPeers failed:",
					err,
				)

				continue
			}

			for _, peer := range storedPeers {

				status := "offline"

				if time.Now().Unix()-peer.LastSeen < 30 {
					status = "online"
				}
				fmt.Printf(
					"\n%s (%s)\n",
					peer.Name,
					peer.ID,
				)

				fmt.Printf(
					"Address: %s\n",
					peer.Address,
				)
				fmt.Printf(
					"Trusted: %t\n",
					peer.Trusted,
				)
				fmt.Printf(
					"Trusts Me: %t\n", peer.TrustsMe,
				)
				info, err := commands.Metadata(
					protocol.PeerInfo{
						ID:      peer.ID,
						Name:    peer.Name,
						Address: peer.Address,
					},
					nodeID,
				)
				fmt.Printf("Known Peers   : %d\n", info.TotalContacts)
				fmt.Printf("Trust Ratio   : %.2f%%\n", info.TrustRatio)
				fmt.Printf("Confidence Score : %.2f\n", info.ConfidenceScore)
				fmt.Printf("Confidence       : %s\n", info.Confidence)

				if err != nil {
					fmt.Println("Metadata unavailable:", err)
				} else {
					fmt.Printf("Trusted By: %d\n", info.TrustedBy)

					fmt.Println("Mutual Trusted:")
					if len(info.MutualTrusted) == 0 {
						fmt.Println("  None")
					} else {
						for _, name := range info.MutualTrusted {
							fmt.Printf("  - %s\n", name)
						}
					}
				}

				if peer.TrustedSince != 0 {

					days := int(
						time.Since(
							time.Unix(
								peer.TrustedSince,
								0,
							),
						).Hours() / 24,
					)

					fmt.Printf(
						"Trust Age: %d days\n",
						days,
					)
				}

				fmt.Printf(
					"Status: %s\n",
					status,
				)
			}

			continue
		}

		if strings.HasPrefix(text, "/use ") {

			target := strings.TrimPrefix(text, "/use ")

			currentTarget = commands.Use(manager, target)

			continue
		}

		if strings.HasPrefix(
			text,
			"/trust ",
		) {

			targetID := strings.TrimPrefix(
				text,
				"/trust ",
			)

			commands.Trust(
				targetID,
				nodeID,
				manager,
			)

			continue
		}

		if strings.HasPrefix(
			text,
			"/untrust ",
		) {

			targetID := strings.TrimPrefix(
				text,
				"/untrust ",
			)

			commands.Untrust(
				targetID,
			)

			continue
		}

		if strings.HasPrefix(text, "/history ") {

			peerID := strings.TrimPrefix(
				text,
				"/history ",
			)

			commands.History(peerID)

			continue
		}

		msg := protocol.Message{
			Type:      protocol.MessageTypeChat,
			SenderID:  nodeID,
			Payload:   text,
			Timestamp: time.Now().Unix(),
		}

		if currentTarget == "" {

			fmt.Println("No active chat target")
			continue
		}

		p := manager.GetPeerByID(currentTarget)

		if p == nil {

			fmt.Println("Target peer not found")
			continue
		}

		chatMsg := data.ChatMessage{
			SenderID: nodeID,
			Text:     text,
			Time:     msg.Timestamp,
			Outgoing: true,
		}

		err := data.SaveMessage(
			currentTarget,
			chatMsg,
		)

		if err != nil {
			fmt.Println("Save failed:", err)
		}

		network.SendMessage(p.Conn, msg)
	}
}
