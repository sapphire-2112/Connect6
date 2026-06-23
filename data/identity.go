package data
import (
	"fmt"
	"time"
	"encoding/json"
	"os"
)

type Identity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	TrustedBy int `json:"trusted_by"`
}


func LoadIdentity() (*Identity, error) {

	data, err := os.ReadFile("data/identity.json")

	if err == nil {
			if len(data) > 0 {

		var identity Identity

		err = json.Unmarshal(data, &identity)
		if err != nil {
			return nil, err
		}

		return &identity, nil
	}
	}
	
	id := fmt.Sprintf(
		"con6-%d",
		time.Now().UnixNano(),
	)

	identity := &Identity{
		ID: id,
	}

	err = SaveIdentity(identity)
	if err != nil {
		return nil, err
	}

	return identity, nil
}


func SaveIdentity(identity *Identity) error {

	data, err := json.MarshalIndent(
		identity,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		"data/identity.json",
		data,
		0644,
	)
}
