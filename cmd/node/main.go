package main

import (
	"bufio"
	"connect6/network"
	"connect6/peer"
	"connect6/protocol"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var manager = peer.NewManager()


func handleConnection(conn net.Conn) {

	p := &peer.Peer{
		ID:      conn.RemoteAddr().String(),
		Address: conn.RemoteAddr().String(),
		Conn:    conn,
	}


	manager.AddPeer(p)

	fmt.Println("Connected:", p.Address)

	go network.ReceiveMessages(conn, func(msg protocol.Message) {

		fmt.Printf(
			"\n[%s] %s\n",
			msg.SenderID,
			msg.Payload,
		)

		fmt.Print("You: ")
	})
}
	

func main() {

		var nodeID string

		fmt.Print("Enter node ID: ")
		fmt.Scanln(&nodeID)

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

	
	

	
	input := bufio.NewScanner(os.Stdin)

for {

	fmt.Print("> ")

	if !input.Scan() {
		return
	}

	text := input.Text()

	
	if strings.HasPrefix(text, "/connect ") {

		address := strings.TrimPrefix(text, "/connect ")

		conn, err := network.Connect(address)
		if err != nil {
			fmt.Println("Connection failed")
			continue
		}

		go handleConnection(conn)

		fmt.Println("Connected to", address)

		continue
	}

	// SHOW PEERS
	if text == "/peers" {

		peers := manager.GetPeers()

		for _, p := range peers {
			fmt.Println(p.Address)
		}

		continue
	}

	
	msg := protocol.Message{
		Type:      protocol.MessageTypeChat,
		SenderID:  nodeID,
		Payload:   text,
		Timestamp: time.Now().Unix(),
	}

	peers := manager.GetPeers()

	for _, p := range peers {
		network.SendMessage(p.Conn, msg)
	}
}
}