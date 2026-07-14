package mobile

import (
	"connect6/data"
	"time"
)

type PeerSummary struct {
	ID       string
	Name     string
	Online   bool
	LastSeen int64
	Trusted  bool
	TrustsMe bool
}

func (e *Engine) GetPeers() ([]PeerSummary, error) {

	storedPeers, err := data.LoadPeers()
	if err != nil {
		return nil, err
	}

	var peers []PeerSummary

	for _, peer := range storedPeers {

		online := time.Now().Unix()-peer.LastSeen < 30

		peers = append(peers, PeerSummary{
			ID:       peer.ID,
			Name:     peer.Name,
			Online:   online,
			LastSeen: peer.LastSeen,
			Trusted:  peer.Trusted,
			TrustsMe: peer.TrustsMe,
		})
	}

	return peers, nil
}