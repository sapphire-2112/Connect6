package peer

import "sync"

type Manager struct {
	Peers map[string]*Peer
	Mutex sync.Mutex
}

func NewManager() *Manager {

	return &Manager{
		Peers: make(map[string]*Peer),
	}
}

func (m *Manager) AddPeer(peer *Peer) {

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.Peers[peer.ID] = peer
}

func (m *Manager) GetPeers() []*Peer {

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	var peers []*Peer

	for _, peer := range m.Peers {
		peers = append(peers, peer)
	}

	return peers
}