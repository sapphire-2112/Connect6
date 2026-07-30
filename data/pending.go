package data
import(
    "os"
    "encoding/json"
    "fmt"
)

type PendingMessage struct {
	ID         string
	SenderID   string
	ReceiverID string
	Content    string
	Timestamp  int64
}


func LoadPendingMessages() ([]PendingMessage, error) {
    var messages []PendingMessage
	 filePath := PendingRequestsPath()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []PendingMessage{}, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return messages, err
	}

	err = json.Unmarshal(data, &messages)
	if err != nil {
		return messages, err
	}

	return messages, nil
}

func SavePendingMessages(messages []PendingMessage) error {

	filePath := PendingRequestsPath()

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

func AddPendingMessage(message PendingMessage) error {
   

    messages, err := LoadPendingMessages()
    if err != nil {
        return err
    }

    fmt.Println("Loaded", len(messages), "pending messages")

    messages = append(messages, message)

    fmt.Println("Saving", len(messages), "messages")

    return SavePendingMessages(messages)
}
func RemovePendingMessage(id string) error {
    messages, err := LoadPendingMessages()
    if err != nil {
        return err
    }

    for i, message := range messages {
        if message.ID == id {
            messages = append(messages[:i], messages[i+1:]...)
            break
        }
    }

    return SavePendingMessages(messages)
}