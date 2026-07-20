package mobile

import (
	"connect6/data"
	"connect6/network"
	"connect6/protocol"
	"connect6/commands"
	"errors"
	"time"
)

func (e *Engine) SendRequest(targetID string) error {

    peerInfo, exists := e.introducedPeers[targetID]
    if !exists {
        return errors.New("peer not known")
    }

    conn, err := commands.Connect(peerInfo.Address)
    if err != nil {
        return err
    }
    defer conn.Close()

    storedPeers, err := data.LoadPeers()
    if err != nil {
        return err
    }

    msg := protocol.Message{
        Type:     protocol.MessageTypeConnectionRequest,
        SenderID: e.identity.ID,
        PeerInfo: &protocol.PeerInfo{
            ID:             e.identity.ID,
            Name:           e.identity.Name,
            Address:        e.identity.Address,
            TrustedBy:      e.identity.TrustedBy,
            TrustedByPeers: e.identity.TrustedByPeers,
            TotalContacts:  len(storedPeers),
        },
        Timestamp: time.Now().Unix(),
    }

    return network.SendMessage(conn, msg)
}