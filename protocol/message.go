package protocol

type Message struct {
	Type      string   `json:"type"`
	SenderID  string   `json:"sender_id"`
	Payload   string   `json:"payload"`
	Peers     []string `json:"peers"`
	Timestamp int64    `json:"timestamp"`
}