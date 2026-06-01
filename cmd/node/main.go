package main

import (
	"bufio"
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"connect6/commands"
	"connect6/data"
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

					fmt.Println(
						"Identity learned:",
						p.ID,
					)
					manager.AddPeer(p)

					return
				}

		if msg.Type == protocol.MessageTypeHeartbeat {

				p.LastSeen = time.Now().Unix()
				p.Online = true
				return
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
					"%s -> %s\n",
					p.ID,
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

		if !p.Connected {

			fmt.Println("Target peer disconnected")
			continue
		}

		network.SendMessage(p.Conn, msg)
}
}