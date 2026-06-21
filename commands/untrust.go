package commands

import (
	"connect6/data"
	"fmt"
)

func Untrust(targetID string) {

	peers, err := data.LoadPeers()

	if err != nil {
		fmt.Println("LoadPeers failed")
		return
	}

	found := false

	for i := range peers {

		if peers[i].ID == targetID {

			peers[i].Trusted = false
			peers[i].TrustedSince = 0

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

	fmt.Println("Untrusted:", targetID)
}