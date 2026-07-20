package mobile

import (
	"connect6/data"
	"connect6/network"
	"connect6/protocol"
	"errors"
	"fmt"
	"time"
)

func (e *Engine) SendMessage(peerID, text string) error {

	if peerID == "" {
		return errors.New("no target peer")
	}

	p := e.manager.GetPeerByID(peerID)
	if p == nil {
		return errors.New("target peer not found")
	}

	msg := protocol.Message{
		Type:      protocol.MessageTypeChat,
		SenderID:  e.identity.ID,
		Payload:   text,
		Timestamp: time.Now().Unix(),
	}

	// Save outgoing message
	chatMsg := data.ChatMessage{
		SenderID: e.identity.ID,
		Text:     text,
		Time:     msg.Timestamp,
		Outgoing: true,
	}

	if err := data.SaveMessage(peerID, chatMsg); err != nil {
		fmt.Println("Save failed:", err)
	}

	// Peer offline -> queue message
	if p.Conn == nil {

		fmt.Println("Connection already closed. Queueing message...")

		pending := data.PendingMessage{
			ID:         fmt.Sprintf("msg-%d", time.Now().UnixNano()),
			SenderID:   e.identity.ID,
			ReceiverID: peerID,
			Content:    text,
			Timestamp:  msg.Timestamp,
		}

		if err := data.AddPendingMessage(pending); err != nil {
			return err
		}

		return nil
	}

	// Try sending
	err := network.SendMessage(p.Conn, msg)
	if err != nil {

		fmt.Println("Connection lost.")

		p.Online = false
		p.Conn = nil

		pending := data.PendingMessage{
			ID:         fmt.Sprintf("msg-%d", time.Now().UnixNano()),
			SenderID:   e.identity.ID,
			ReceiverID: peerID,
			Content:    text,
			Timestamp:  msg.Timestamp,
		}

		if err := data.AddPendingMessage(pending); err != nil {
			return err
		}

		return nil
	}

	return nil
}