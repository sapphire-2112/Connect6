package main

import (
	"bufio"
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"connect6/commands"
	"connect6/data"
	"connect6/crypto"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	
)

var manager = peer.NewManager()
var currentTarget string
var nodeID string


func handleConnection(conn net.Conn) {

	p := &peer.Peer{
		ID:      "",
		Address: conn.RemoteAddr().String(),
		Connected: true,
		Conn:    conn,
		Online:    true,
		LastSeen:  time.Now().Unix(),
	}
	


	

	fmt.Println("Connected:", p.Address)


	join := protocol.Message{
	Type:      protocol.MessageTypeJoin,
	SenderID:  nodeID,
	Timestamp: time.Now().Unix(),
		}

		network.SendMessage(conn, join)

	go network.ReceiveMessages(conn, func(msg protocol.Message) {


			if msg.Type == protocol.MessageTypeJoin {

					if p.ID != "" {
						return
					}

					p.ID = msg.SenderID
					manager.AddPeer(p)

					fmt.Println(
						"Identity learned:",
						p.ID,
					)
					p.Online = true

					host, _, err := net.SplitHostPort(p.Address)
						if err != nil {
							fmt.Println("Address parse failed:", err)
							return
						}

						realAddress := net.JoinHostPort(host, "8080")

					

					storedPeer := data.StoredPeer{
						ID:       p.ID,
						Address:  realAddress,
						Trusted:  false,
						LastSeen: p.LastSeen,
					}

					peers, _ := data.LoadPeers()

					peers = append(peers, storedPeer)

					err = data.SavePeers(peers)
					if err != nil {
						fmt.Println("SavePeers failed:", err)
					}

					return
				}
		if msg.Type == protocol.MessageTypeHeartbeat {

				p.LastSeen = time.Now().Unix()
				p.Online = true
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

		identity, err := data.LoadIdentity()
			if err != nil {
				fmt.Println(err)
				return
			}

			nodeID = identity.ID
			fmt.Println("Node ID:", nodeID)

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

					msg := protocol.Message{
						Type:      protocol.MessageTypeHeartbeat,
						SenderID:  nodeID,
						Timestamp: time.Now().Unix(),
					}


					peers := manager.GetPeers()

					for _, p := range peers {

						if p.Connected {

							network.SendMessage(p.Conn, msg)
						}
					}
				}
			}()

			go func() {

				for {

					time.Sleep(5 * time.Second)

					peers := manager.GetPeers()

					for _, p := range peers {

						if time.Now().Unix()-p.LastSeen > 30 {

							p.Online = false
							  if p.Conn != nil {
									p.Conn.Close()
								}
							}
					}
				}
			}()
	
	

	
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


	
	if strings.HasPrefix(text, "/connect ") {

		address := strings.TrimPrefix(text, "/connect ")

		conn, err := commands.Connect(address)
		if err != nil {
			fmt.Println("Connection failed")
			continue
		}
		join := protocol.Message{
		Type:      protocol.MessageTypeJoin,
		SenderID:  nodeID,
		Timestamp: time.Now().Unix(),
	}

		network.SendMessage(conn, join)

		go handleConnection(conn)

		fmt.Println("Connected to", address)

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

		peers := manager.GetPeers()

		for _, p := range peers {
			status := "offline"

				if p.Online {
					status = "online"
				}

				fmt.Printf(
					" %s -> %s (%s)\n",
					p.ID,
					status,
					p.Address,
				)
		}

		continue
	}

	if strings.HasPrefix(text, "/use ") {

	target := strings.TrimPrefix(text, "/use ")

	currentTarget = commands.Use(manager, target)

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