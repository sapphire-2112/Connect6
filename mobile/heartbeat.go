package mobile

import (
	"connect6/data"
	"connect6/protocol"
	"connect6/network"
	"time"
	"connect6/commands"
)

func (e *Engine) heartbeatLoop() {

    for {

        time.Sleep(10 * time.Second)

        peers, err := data.LoadPeers()
        if err != nil {
            continue
        }

        for _, peer := range peers {

            conn, err := commands.Connect(peer.Address)
            if err != nil {
                continue
            }

            msg := protocol.Message{
                Type:      protocol.MessageTypeHeartbeat,
                SenderID:  e.identity.ID,
                Timestamp: time.Now().Unix(),
            }

            network.SendMessage(conn, msg)

            conn.Close()
        }
    }
}