package mobile

import (
	"connect6/network"
	"connect6/protocol"
	"fmt"
	"time"
)

func (e *Engine) RejectRequest(peerID string) error {

	e.mu.Lock()

	_, ok := e.pendingRequests[peerID]
	if !ok {
		e.mu.Unlock()
		return fmt.Errorf("pending request not found")
	}

	delete(e.pendingRequests, peerID)

	e.mu.Unlock()

	p := e.manager.GetPeerByID(peerID)
	if p == nil {
		return fmt.Errorf("peer not connected")
	}

	msg := protocol.Message{
		Type:      protocol.MessageTypeConnectionReject,
		SenderID:  e.identity.ID,
		Timestamp: time.Now().Unix(),
	}

	if err := network.SendMessage(p.Conn, msg); err != nil {
		return err
	}

	return nil
}