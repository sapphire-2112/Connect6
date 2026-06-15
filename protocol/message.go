package protocol

type Message struct {
	Type       string      `json:"type"`
	SenderID   string      `json:"sender_id"`
	Payload    string      `json:"payload"`
	Peers      []string    `json:"peers"`
	PeerInfos  []PeerInfo  `json:"peer_infos"`
	PeerInfo   *PeerInfo   `json:"peer_info,omitempty"`
	Timestamp  int64       `json:"timestamp"`
}