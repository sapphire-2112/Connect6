package commands

import (
	"connect6/data"
	"fmt"
	"time"
)

func Trust(targetID string) {

	peers, err := data.LoadPeers()

	if err != nil {
		fmt.Println("LoadPeers failed")
		return
	}

	found := false

	for i := range peers {

		if peers[i].ID == targetID {

			peers[i].Trusted = true
			peers[i].TrustedSince = time.Now().Unix()

			found = true
			break
		}
	}

	if !found {

		fmt.Println("Peer not found")
		return
	}

	err = data.SavePeers(peers)

	if err != nil {

		fmt.Println("SavePeers failed")
		return
	}

	fmt.Println("Trusted:", targetID)
}