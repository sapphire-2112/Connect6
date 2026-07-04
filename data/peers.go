package data

import (
	"encoding/json"
	"os"
	
)

type StoredPeer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Trusted  bool   `json:"trusted"`
	TrustedSince int64  `json:"trusted_since"`
	TrustedBy int	`json:"trusted_by"`
	TrustsMe bool `json:"trusts_me"`
	LastSeen int64  `json:"last_seen"`
}

func SavePeers(peers []StoredPeer) error {

	filePath := "data/peers.json"
	data, err := json.MarshalIndent(peers, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

	

func LoadPeers() ([]StoredPeer, error) {

	var peers []StoredPeer
	if _, err := os.Stat("data/peers.json"); os.IsNotExist(err) {
		return []StoredPeer{}, nil
	}

	data, err := os.ReadFile("data/peers.json")
	if err != nil {
		return peers, err
	}

	err = json.Unmarshal(data, &peers)
	if err != nil {
		return peers, err
	}

	return peers, nil
}
func GetTrustedByCount() int {

    peers, err := LoadPeers()

    if err != nil {
        return 0
    }

    count := 0

    for _, p := range peers {

        if p.TrustsMe {

            count++
        }
    }

    return count
}