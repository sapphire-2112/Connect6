package mobile

import (
	"connect6/data"
	"connect6/network"
	"connect6/protocol"
	"errors"
	"fmt"
	"time"
)

func (e *Engine) Trust(peerID string) error {

	peers, err := data.LoadPeers()
	if err != nil {
		return err
	}

	found := false

	for i := range peers {

		if peers[i].ID == peerID {

			peers[i].Trusted = true
			peers[i].TrustedSince = time.Now().Unix()

			found = true
			break
		}
	}

	if !found {
		return errors.New("peer not found")
	}

	if err := data.SavePeers(peers); err != nil {
		return err
	}

	p := e.manager.GetPeerByID(peerID)

	if p != nil && p.Conn != nil {

		msg := protocol.Message{
			Type:      protocol.MessageTypeTrust,
			SenderID:  e.identity.ID,
			Timestamp: time.Now().Unix(),
		}

		network.SendMessage(p.Conn, msg)
	}

	fmt.Println("Trusted:", peerID)

	return nil
}