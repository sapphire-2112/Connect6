# Connect6 Development Plan

## Current Stage

Connect6 currently functions as a decentralized identity-based peer-to-peer communication runtime where independently operating nodes communicate directly over IPv6 TCP connections.

Implemented capabilities include:

* concurrent peer handling
* active peer sessions
* heartbeat presence tracking
* online/offline peer detection
* targeted messaging
* persistent node identities
* identity handshake protocol
* persistent peer storage
* automatic peer loading
* automatic reconnect
* local chat persistence
* conversation history retrieval
* peer discovery protocol (`/peersof`)
* decentralized node runtime behavior
* Docker-based distributed testing

---

# Core Architectural Goals

## Decentralized Identity

Move from transport-based addressing toward persistent node identities that survive network changes.

### Implemented

* persistent node identity generation
* identity storage (`identity.json`)
* identity handshake between peers

### Planned

* cryptographic identities
* public/private key pairs
* cryptographic peer verification
* signed peer introductions

---

## Persistent Peer Graph

Build a decentralized communication graph where:

* peers persist locally
* relationships survive restarts
* peer discovery expands network visibility
* future trust relationships become possible

### Implemented

* peer persistence (`peers.json`)
* peer loading on startup
* automatic reconnect
* peer discovery (`/peersof`)

### Planned

* trust metadata
* peer reputation metadata
* trusted peer relationships
* decentralized trust graph

---

## Secure Communication Layer

Introduce encrypted peer-to-peer communication.

### Planned

* end-to-end encryption
* encrypted local vault
* encrypted protocol payloads
* secure session establishment
* encrypted chat storage
* message signing

---

## Presence & Session Management

Expand runtime communication behavior.

### Implemented

* heartbeat protocol
* online/offline tracking
* last-seen monitoring
* automatic reconnect

### Planned

* session deduplication
* heartbeat optimization
* identity-based session recovery

---

## Messaging System Evolution

Improve messaging architecture.

### Implemented

* direct peer messaging
* active conversation targeting
* local chat persistence
* conversation history retrieval

### Planned

* connection requests
* peer acceptance/rejection workflow
* peer introductions
* multi-conversation handling
* delivery acknowledgements
* offline message handling
* message synchronization

---

## Decentralized Networking Goals

Long-term networking direction includes:

* peer discovery
* peer introductions
* connection request routing
* trust-based onboarding
* peer-of-peer visibility
* distributed communication graph expansion
* resilient communication routing

---

## Local Ownership Philosophy

Connect6 prioritizes:

* user-owned identity
* local-first storage
* decentralized trust
* metadata minimization
* reduced centralized dependency
* user-controlled communication

---

## Long-Term Research Areas

The project serves as an exploration into:

* distributed systems
* decentralized networking
* protocol engineering
* concurrent runtime design
* cryptographic communication systems
* peer-to-peer architectures
* resilient communication ecosystems

---

# Current Priorities

1. Peer Introduction System
2. Connection Request Workflow
3. Peer Acceptance / Rejection System
4. Session Deduplication
5. Cryptographic Identity Verification
6. Message Signing
7. Encrypted Local Storage
8. Trust Graph Foundation
9. Group Communication Support

---

# Current Milestones

## Completed

✓ IPv6 networking

✓ Multi-peer communication

✓ Heartbeat presence system

✓ Identity persistence

✓ Identity handshake

✓ Direct messaging

✓ Chat persistence

✓ Conversation history retrieval

✓ Persistent peer storage

✓ Automatic peer loading

✓ Automatic reconnect

✓ Peer discovery protocol (`/peersof`)

---

## Next Milestone

→ Peer Introduction & Connection Request System

### Phase 1

✓ Peer Discovery (`/peersof`)

### Phase 2

→ Connection Request Protocol

* request connection to discovered peer
* route request through mutual peer
* display requester metadata

### Phase 3

→ Acceptance Workflow

* accept request
* reject request
* establish trusted relationship

### Phase 4

→ Trust Metadata

* introduced by
* mutual peers
* trust level
* peer notes

### Phase 5

→ Trust Graph Foundation

* trusted peers
* introduced peers
* trust propagation research
