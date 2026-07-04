package protocol

const (
	MessageTypeChat     = "chat"
	MessageTypePeerList = "peer_list"
	MessageTypeJoin     = "join"
	MessageTypeHeartbeat = "heartbeat"
	MessageTypePeerListRequest  = "peer_list_request"
	MessageTypePeerListResponse = "peer_list_response"
	MessageTypeConnectionRequest = "connection_request"
	MessageTypeConnectionAccept  = "connection_accept"
	MessageTypeConnectionReject  = "connection_reject"
	MessageTypeTrust   = "trust"
	MessageTypeUntrust = "untrust"
	MessageTypeReputationUpdate = "reputation_update"
	MessageTypeMetadataRequest  = "metadata_request"
	MessageTypeMetadataResponse = "metadata_response"
)