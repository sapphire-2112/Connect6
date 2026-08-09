package mobile

import (
	"connect6/data"
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

    for _, p := range storedPeers {
        if p.ID == peer.ID {
            return nil
        }
    }

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

    return data.SavePeers(storedPeers)
}