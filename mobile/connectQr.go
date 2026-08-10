package mobile

import (
	"connect6/commands"
	"connect6/data"
	"fmt"
	"time"
)

func (e *Engine) ConnectFromQR(payload string) error {

	peer, err := ParseQRCodePayload(payload)
	if err != nil {
		return err
	}

	storedPeers, err := data.LoadPeers()
	if err != nil {
		return err
	}

	found := false

	for _, p := range storedPeers {
		if p.ID == peer.ID {
			found = true
			break
		}
	}

	if !found {

		storedPeers = append(storedPeers, data.StoredPeer{
			ID:           peer.ID,
			Name:         peer.Name,
			Address:      peer.Address,
			Trusted:      true,
			TrustedSince: time.Now().Unix(),
			TrustedBy:    0,
			TrustsMe:     true,
			LastSeen:     0,
		})

		err = data.SavePeers(storedPeers)
		if err != nil {
			return err
		}

		fmt.Println("Saved peer:", peer.Name)
	}

	fmt.Println("Connecting to:", peer.Address)

	conn, err := commands.Connect(peer.Address)
	if err != nil {
		return err
	}

	go e.handleConnection(conn)

	fmt.Println("Connection initiated")

	return nil
}