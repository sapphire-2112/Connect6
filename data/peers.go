package data
import (
	"encoding/json"
	"os"
)

type StoredPeer struct {
	ID        string `json:"id"`
	Address   string `json:"address"`
	Trusted   bool   `json:"trusted"`
	LastSeen  int64  `json:"last_seen"`
}
func SavePeers(peers []StoredPeer) error {

	data, err := json.MarshalIndent(peers, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(
		"data/peers.json",
		data,
		0644,
	)
}

func LoadPeers() ([]StoredPeer, error) {

	var peers []StoredPeer

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