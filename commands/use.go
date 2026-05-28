package commands

import (
	"connect6/peer"
	"fmt"
)

func Use(
	manager *peer.Manager,
	target string,
) string {

	p := manager.GetPeerByID(target)

	if p == nil {

		fmt.Println("Peer not found")
		return ""
	}

	fmt.Println("Now chatting with", target)

	return target
}