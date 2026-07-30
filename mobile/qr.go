package mobile

import (
    "encoding/json"
    "errors"
)

type QRPeer struct {
    Type    string `json:"type"`
    Version int    `json:"version"`
    ID      string `json:"id"`
    Name    string `json:"name"`
    Address string `json:"address"`
}

func (e *Engine) GetQRCodePayload() (string, error) {
    if e.identity == nil {
        return "", errors.New("engine not initialized")
    }

    qr := QRPeer{
        Type:    "peer",
        Version: 1,
        ID:      e.identity.ID,
        Name:    e.identity.Name,
        Address: e.identity.Address,
    }

    b, err := json.Marshal(qr)
    if err != nil {
        return "", err
    }

    return string(b), nil
}