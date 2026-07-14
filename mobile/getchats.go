package mobile
import (
	"connect6/data"
)
func (e *Engine) GetChats(peerID string) ([]data.ChatMessage, error) {
    return data.LoadChat(peerID)
}