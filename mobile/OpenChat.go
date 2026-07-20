package mobile

import (
	"connect6/network"
	"connect6/protocol"
	"errors"
	"time"
)

func (e *Engine) OpenChat(peerID string) error {

	p := e.manager.GetPeerByID(peerID)
	if p == nil {
		return errors.New("peer not found")
	}

	if p.Conn == nil {
		return nil
	}

	msg := protocol.Message{
		Type:      protocol.MessageTypePendingSyncRequest,
		SenderID:  e.identity.ID,
		Timestamp: time.Now().Unix(),
	}

	return network.SendMessage(p.Conn, msg)
}