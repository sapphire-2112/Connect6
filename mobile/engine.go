package mobile

import (
	"connect6/data"
	"connect6/peer"
	"connect6/protocol"
	"connect6/network"
    "connect6/crypto"
    "fmt"
    "path/filepath"
    "os"
	"net"
    "sync"
)

type Engine struct {
	manager *peer.Manager

	identity *data.Identity

	knownPeers []data.StoredPeer
    storagePath string

	listener net.Listener
     advertisedAddress string

	introducedPeers map[string]protocol.PeerInfo
	pendingRequests map[string]protocol.PeerInfo
    mu sync.RWMutex
}

func NewEngine() *Engine {
	return &Engine{
		manager:          peer.NewManager(),
		introducedPeers:  make(map[string]protocol.PeerInfo),
		pendingRequests:  make(map[string]protocol.PeerInfo),
	}
}

func (e *Engine) StartNode() error {

    identity, err := data.LoadIdentity()
    if err != nil {
        return err
    }
    e.identity = identity
   if e.advertisedAddress != "" &&
    e.identity.Address != e.advertisedAddress {

    e.identity.Address = e.advertisedAddress

    if err := data.SaveIdentity(e.identity); err != nil {
        return err
    }
}

    knownPeers, err := data.LoadPeers()
    if err != nil {
        return err
    }
    e.knownPeers = knownPeers

    os.MkdirAll(filepath.Dir(data.CertPath()), 0755)

        if _, err := os.Stat(data.CertPath()); os.IsNotExist(err) {

            err := crypto.GenerateCertificate(
                data.CertPath(),
                data.KeyPath(),
            )
            if err != nil {
                    return err
                }

                fmt.Println("Generated TLS certificate")
            }
            

     

        listener, err := network.StartListener("[::]:8080")
        if err != nil {
            return err
        }
        e.listener = listener
        fmt.Println("Advertised:", e.identity.Address)
        fmt.Println("Listening :", "[::]:8080")

        go e.acceptLoop()
        go e.heartbeatLoop()
        go e.reconnectLoop()

        return nil
    }

    func (e *Engine) Version() string {
        return "1.0"
    }
func (e *Engine) SetStoragePath(path string) {
	data.SetStoragePath(path)
}
func (e *Engine) SetAdvertisedAddress(address string) {
    e.advertisedAddress = address
}
// func (e *Engine) StopNode()

// // --------------------
// // Connections
// // --------------------


// func (e *Engine) Disconnect(peerID string) error

// // --------------------
// // Messaging
// // --------------------



// func (e *Engine) DeleteChat(peerID string) error

// // --------------------
// // Requests
// // --------------------



// // --------------------
// // QR
// // --------------------

// func (e *Engine) GenerateInviteQR()

// func (e *Engine) ParseInviteQR(contents string) error

// // --------------------
// // Identity
// // --------------------



// func (e *Engine) Rename(name string) error

// // --------------------
// // Offline
// // --------------------

// func (e *Engine) SyncPendingMessages(peerID string)

// func (e *Engine) PendingMessageCount() int