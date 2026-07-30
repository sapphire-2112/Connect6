package mobile

import (
    "connect6/data"
    "encoding/json"
)

func (e *Engine) GetChatsJSON(peerID string) (string, error) {
    chats, err := data.LoadChat(peerID)
    if err != nil {
        return "", err
    }

    b, err := json.Marshal(chats)
    if err != nil {
        return "", err
    }

    return string(b), nil
}