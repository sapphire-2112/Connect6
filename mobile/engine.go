package mobile

import (
	"connect6/data"
	"connect6/peer"
	"connect6/protocol"
	"connect6/network"
	"net"
    "sync"
)

type Engine struct {
	manager *peer.Manager

	identity *data.Identity

	knownPeers []data.StoredPeer

	listener net.Listener

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

    knownPeers, err := data.LoadPeers()
    if err != nil {
        return err
    }
    e.knownPeers = knownPeers

    listener, err := network.StartListener(e.identity.Address)
    if err != nil {
        return err
    }
    e.listener = listener

    go e.acceptLoop()
    go e.heartbeatLoop()
    go e.reconnectLoop()

    return nil
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

// // Send a connection request
// func (e *Engine) SendRequest(peerID string) error

// // All incoming pending requests
// func (e *Engine) GetPendingRequests()

// // Accept / Reject
// func (e *Engine) AcceptRequest(peerID string) error

// func (e *Engine) RejectRequest(peerID string) error





// func (e *Engine) GetPeersOf(peerID string)

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