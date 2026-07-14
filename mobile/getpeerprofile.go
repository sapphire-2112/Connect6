package mobile

import (
	"connect6/commands"
	"connect6/data"
	"connect6/protocol"
	"fmt"
	"time"
)

type PeerProfile struct {
	ID              string
	Name            string
	Online          bool
	TrustedBy       int
	TotalContacts   int
	TrustRatio      float64
	ConfidenceScore float64
	Confidence      string
	MutualTrusted   []string
}

func (e *Engine) GetPeerProfile(peerID string) (PeerProfile, error) {

	peers, err := data.LoadPeers()
	if err != nil {
		return PeerProfile{}, err
	}

	var peer data.StoredPeer
	found := false

	for _, p := range peers {
		if p.ID == peerID {
			peer = p
			found = true
			break
		}
	}

	if !found {
		return PeerProfile{}, fmt.Errorf("peer not found")
	}

	identity, err := data.LoadIdentity()
	if err != nil {
		return PeerProfile{}, err
	}

	info, err := commands.Metadata(
		protocol.PeerInfo{
			ID:      peer.ID,
			Name:    peer.Name,
			Address: peer.Address,
		},
		identity.ID,
	)
	if err != nil {
		return PeerProfile{}, err
	}

	online := false
	if time.Now().Unix()-peer.LastSeen < 30 {
		online = true
	}

	profile := PeerProfile{
		ID:              peer.ID,
		Name:            peer.Name,
		Online:          online,
		TrustedBy:       info.TrustedBy,
		TotalContacts:   info.TotalContacts,
		TrustRatio:      info.TrustRatio,
		ConfidenceScore: info.ConfidenceScore,
		Confidence:      info.Confidence,
		MutualTrusted:   info.MutualTrusted,
	}

	return profile, nil
}