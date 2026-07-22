package mobile

import (
	"connect6/network"
	"connect6/protocol"
	"fmt"
	"time"
	"connect6/data"
)

func (e *Engine) AcceptRequest(peerID string) error {

    e.mu.Lock()

    _, ok := e.pendingRequests[peerID]
    if !ok {
        e.mu.Unlock()
        return fmt.Errorf("pending request not found")
    }

    e.mu.Unlock()

    p := e.manager.GetPeerByID(peerID)
    if p == nil {
        return fmt.Errorf("peer not connected")
    }

    msg := protocol.Message{
        Type:      protocol.MessageTypeConnectionAccept,
        SenderID:  e.identity.ID,
        Timestamp: time.Now().Unix(),
    }

    if err := network.SendMessage(p.Conn, msg); err != nil {
        return err
    }

    e.mu.Lock()
    defer e.mu.Unlock()

    // Avoid duplicate peers
    for _, peer := range e.knownPeers {
        if peer.ID == peerID {
            delete(e.pendingRequests, peerID)
            return nil
        }
    }

    stored := data.StoredPeer{
        ID:           p.ID,
        Name:         p.Name,
        Address:      p.Address,
        Trusted:      false,
        TrustedSince: 0,
        TrustedBy:    0,
        TrustsMe:     false,
        LastSeen:     time.Now().Unix(),
    }

    e.knownPeers = append(e.knownPeers, stored)

    if err := data.SavePeers(e.knownPeers); err != nil {
        return err
    }

    delete(e.pendingRequests, peerID)

    return nil
}