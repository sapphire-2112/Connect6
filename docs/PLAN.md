# Connect6 Development Plan

## Current Stage

Connect6 currently functions as a decentralized identity-based peer-to-peer communication runtime where independently operating nodes communicate directly over IPv6 TCP connections.

Implemented capabilities include:

* concurrent peer handling
* active peer sessions
* heartbeat presence tracking
* decentralized online/offline detection
* targeted messaging
* persistent node identities
* identity handshake protocol
* persistent peer storage
* automatic peer loading
* automatic reconnect
* local chat persistence
* conversation history retrieval
* peer discovery protocol (`/peersof`)
* peer introduction workflow
* connection request workflow
* peer acceptance workflow
* pending request management
* persistent peer relationships
* decentralized node runtime behavior
* Docker-based distributed testing

---

# Core Architectural Goals

## Decentralized Identity

Move from transport-based addressing toward persistent node identities that survive network changes.

### Implemented

* persistent node identity generation
* identity storage (`identity.json`)
* identity handshake protocol
* identity exchange during connection establishment

### Planned

* cryptographic identities
* public/private key pairs
* cryptographic peer verification
* signed peer introductions
* identity reputation metadata

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
* peer introductions
* connection requests
* peer acceptance workflow
* persistent peer relationships

### Planned

* trust metadata
* peer reputation metadata
* trusted peer relationships
* decentralized trust graph
* trust propagation research

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
* decentralized presence detection
* last-seen monitoring
* automatic reconnect

### Planned

* session deduplication
* heartbeat optimization
* identity-based session recovery
* connection lifecycle optimization

---

## Messaging System Evolution

Improve messaging architecture.

### Implemented

* direct peer messaging
* active conversation targeting
* local chat persistence
* conversation history retrieval
* peer introduction workflow
* connection request workflow
* peer acceptance workflow

### Planned

* multi-conversation handling
* delivery acknowledgements
* offline message handling
* message synchronization
* message forwarding
* media transfer support

---

## Decentralized Networking Goals

Long-term networking direction includes:

* peer discovery
* peer introductions
* trust-based onboarding
* peer-of-peer visibility
* distributed communication graph expansion
* resilient communication routing
* decentralized social graph research

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
* trust graph design
* decentralized identity systems

---

# Current Priorities

1. Trust Graph Foundation
2. Session Deduplication
3. Cryptographic Identity Verification
4. Message Signing
5. End-to-End Encryption
6. Encrypted Local Storage
7. Group Communication Support
8. NAT Traversal
9. Offline Message Synchronization

---

# Current Milestones

## Completed

✓ IPv6 networking

✓ Multi-peer communication

✓ Persistent node identities

✓ Identity handshake protocol

✓ Persistent peer storage

✓ Automatic peer loading

✓ Automatic reconnect

✓ Direct peer messaging

✓ Chat persistence

✓ Conversation history retrieval

✓ Peer discovery protocol (`/peersof`)

✓ Peer introduction workflow

✓ Connection request workflow

✓ Peer acceptance workflow

✓ Persistent peer relationships

✓ Heartbeat-based decentralized presence

✓ Online/offline peer detection

---

## Current Milestone

→ Trust Graph Foundation

### Phase 1 — Discovery

✓ Peer Discovery (`/peersof`)

✓ Introduced Peer Visibility

---

### Phase 2 — Relationship Establishment

✓ Connection Request Workflow

✓ Pending Request Management

✓ Peer Acceptance Workflow

✓ Persistent Peer Relationships

---

### Phase 3 — Presence

✓ Heartbeat Protocol

✓ Last-Seen Tracking

✓ Decentralized Online/Offline Detection

---

### Phase 4 — Trust Metadata

→ In Progress

* trusted peers
* introduced-by tracking
* mutual peer visibility
* trust level metadata
* peer notes

---

### Phase 5 — Trust Graph Foundation

* trust propagation research
* trust scoring models
* decentralized onboarding models
* trusted introduction paths
* social graph evolution
