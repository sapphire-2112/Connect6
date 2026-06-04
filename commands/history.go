package commands

import (
	"connect6/data"
	"fmt"
)

func History(peerID string) {

	messages, err := data.LoadChat(peerID)
	if err != nil {
		fmt.Println("Failed to load history:", err)
		return
	}

	for _, msg := range messages {

		if msg.Outgoing {
			fmt.Println("You:", msg.Text)
		} else {
			fmt.Println(msg.SenderID+":", msg.Text)
		}
	}
}