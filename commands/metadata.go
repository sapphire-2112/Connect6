package commands

import (
	"connect6/data"
	"connect6/network"
	"connect6/protocol"
	"time"
)

type MetadataInfo struct {
    TrustedBy     int
    TotalContacts int
    TrustRatio    float64
    MutualTrusted []string
}

func Metadata(
	peer protocol.PeerInfo,
	nodeID string,
) (MetadataInfo, error) {

	conn, err := Connect(peer.Address)
	if err != nil {
		return MetadataInfo{}, err
	}
	defer conn.Close()

	request := protocol.Message{
		Type:      protocol.MessageTypeMetadataRequest,
		SenderID:  nodeID,
		Timestamp: time.Now().Unix(),
	}

	if err := network.SendMessage(conn, request); err != nil {
		return MetadataInfo{}, err
	}

	msg, err := network.ReceiveMetadataResponse(conn)
	if err != nil {
		return MetadataInfo{}, err
	}

	info := MetadataInfo{
		TrustedBy: msg.PeerInfo.TrustedBy,
		TotalContacts: msg.PeerInfo.TotalContacts,
	}
	info.TrustRatio = CalculateTrustRatio(
    info.TrustedBy,
    info.TotalContacts,
)

	storedPeers, err := data.LoadPeers()
	if err != nil {
		return info, nil
	}

	for _, trustedID := range msg.PeerInfo.TrustedByPeers {
		for _, myPeer := range storedPeers {
			if myPeer.ID == trustedID {
				info.MutualTrusted = append(info.MutualTrusted, myPeer.Name)
			}
		}
	}

	return info, nil
}