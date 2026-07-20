package mobile

import (
	"connect6/commands"
	"fmt"
	"time"
)
func (e *Engine) reconnectLoop() {

    for {

        time.Sleep(15 * time.Second)

        peers := e.manager.GetPeers()

        for _, p := range peers {

            if p.Online {
                continue
            }

            conn, err := commands.Connect(p.Address)
            if err != nil {
                continue
            }

            p.Online = true

            fmt.Println("Reconnected:", p.ID)

            go e.handleConnection(conn)
        }
    }
}