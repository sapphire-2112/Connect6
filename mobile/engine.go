package mobile
import (
)
type Engine struct {
}

func NewEngine() *Engine

// --------------------
// Node Lifecycle
// --------------------

func (e *Engine) StartNode() error

func (e *Engine) StopNode()

// --------------------
// Connections
// --------------------


func (e *Engine) Disconnect(peerID string) error

// --------------------
// Messaging
// --------------------

func (e *Engine) SendMessage(peerID, text string) error


func (e *Engine) DeleteChat(peerID string) error

// --------------------
// Requests
// --------------------

// Send a connection request
func (e *Engine) SendRequest(peerID string) error

// All incoming pending requests
func (e *Engine) GetPendingRequests()

// Accept / Reject
func (e *Engine) AcceptRequest(peerID string) error

func (e *Engine) RejectRequest(peerID string) error

// --------------------
// Trust
// --------------------

func (e *Engine) Trust(peerID string) error

func (e *Engine) Untrust(peerID string) error




func (e *Engine) GetPeersOf(peerID string)

// --------------------
// QR
// --------------------

func (e *Engine) GenerateInviteQR()

func (e *Engine) ParseInviteQR(contents string) error

// --------------------
// Identity
// --------------------



func (e *Engine) Rename(name string) error

// --------------------
// Offline
// --------------------

func (e *Engine) SyncPendingMessages(peerID string)

func (e *Engine) PendingMessageCount() int