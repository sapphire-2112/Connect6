package protocol

type PeerInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	TrustedBy int    `json:"trusted_by"`
	TrustedByPeers []string `json:"trusted_by_peers"`
}