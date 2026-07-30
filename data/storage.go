package data

import "path/filepath"

var storagePath = "data"


func SetStoragePath(path string) {
	storagePath = path
}

func IdentityPath() string {
	return filepath.Join(storagePath, "identity.json")
}

func PeersPath() string {
	return filepath.Join(storagePath, "peers.json")
}

func PendingRequestsPath() string {
	return filepath.Join(storagePath, "pending_requests.json")
}

func ChatsDir() string {
	return filepath.Join(storagePath, "chats")
}
func CertPath() string {
	return filepath.Join(storagePath, "cert.pem")
}

func KeyPath() string {
	return filepath.Join(storagePath, "key.pem")
}