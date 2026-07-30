package data
import (
	"encoding/json"
	"os"
	"path/filepath"
)

type ChatMessage struct {
	SenderID string `json:"sender_id"`
	Text     string `json:"text"`
	Time     int64  `json:"time"`
	Outgoing bool   `json:"outgoing"`
}

func LoadChat(peerID string) ([]ChatMessage, error) {

	filePath := filepath.Join(ChatsDir(), peerID+".json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {

		return []ChatMessage{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var messages []ChatMessage

	err = json.NewDecoder(file).Decode(&messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func SaveMessage(peerID string,msg ChatMessage,) error{
	os.MkdirAll("data/chats", 0755)
	filePath := "data/chats/" + peerID + ".json"

	var messages []ChatMessage
	if _, err := os.Stat(filePath); err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(&messages)
		if err != nil {
			return err
		}
	}

	messages = append(messages, msg)
	data, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}