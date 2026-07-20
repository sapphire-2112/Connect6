package mobile

import "connect6/protocol"

func (e *Engine) GetPendingRequests() []protocol.PeerInfo {
		e.mu.RLock()
	defer e.mu.RUnlock()

	requests := make([]protocol.PeerInfo, 0, len(e.pendingRequests))

	for _, peer := range e.pendingRequests {
		requests = append(requests, peer)
	}

	return requests
}