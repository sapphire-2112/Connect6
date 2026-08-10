package mobile

import (
	"connect6/data"
	"encoding/json"
	"time"
)

type PeerSummary struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Online   bool   `json:"online"`
	LastSeen int64  `json:"last_seen"`
	Trusted  bool   `json:"trusted"`
	TrustsMe bool   `json:"trusts_me"`
}

func (e *Engine) GetPeers() (string, error) {

	storedPeers, err := data.LoadPeers()
	if err != nil {
		return "", err
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

	b, err := json.Marshal(peers)
	if err != nil {
		return "", err
	}

	return string(b), nil
}