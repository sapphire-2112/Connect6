# Connect6 Development Plan

# Current Stage

Connect6 currently functions as a decentralized identity-based peer-to-peer communication runtime where independently operating nodes communicate directly over IPv6 TCP connections.

Implemented capabilities include:

* concurrent peer handling
* active peer sessions
* heartbeat presence tracking
* online/offline peer detection
* targeted messaging
* persistent node identities
* identity handshake protocol
* local chat persistence
* conversation history retrieval
* decentralized node runtime behavior
* Docker-based distributed testing

---

# Core Architectural Goals

## Decentralized Identity

Move from transport-based addressing toward persistent node identities that survive network changes.

Implemented:

* persistent node identity generation
* identity storage (`identity.json`)
* identity handshake between peers

Planned:

* cryptographic identities
* public/private key pairs
* cryptographic peer verification
* signed peer introductions

---

# Persistent Peer Graph

Build a decentralized trust-based communication graph where:

* peers persist locally
* trusted nodes survive restarts
* relationships are user-controlled

Planned:

* peer persistence (`peers.json`)
* automatic peer loading
* peer metadata persistence
* trust management
* peer reputation metadata

---

# Secure Communication Layer

Introduce encrypted peer-to-peer communication.

Planned:

* end-to-end encryption
* encrypted local vault
* encrypted protocol payloads
* secure session establishment
* encrypted chat storage

---

# Presence & Session Management

Expand runtime communication behavior.

Implemented:

* heartbeat protocol
* online/offline tracking
* last-seen monitoring

Planned:

* automatic background reconnect
* session deduplication
* decentralized presence propagation
* heartbeat optimization
* identity-based session recovery

---

# Messaging System Evolution

Improve messaging architecture.

Implemented:

* direct peer messaging
* active conversation targeting
* local chat persistence
* conversation history retrieval

Planned:

* multi-conversation handling
* offline message handling
* peer conversation contexts
* delivery acknowledgements
* message synchronization

---

# Decentralized Networking Goals

Long-term networking direction includes:

* decentralized peer discovery
* invite-based onboarding
* peer introductions
* peer-of-peer visibility
* distributed communication graph expansion
* resilient communication routing

---

# Local Ownership Philosophy

Connect6 prioritizes:

* user-owned identity
* local-first storage
* decentralized trust
* metadata minimization
* reduced centralized dependency
* user-controlled communication

---

# Long-Term Research Areas

The project also serves as exploration into:

* distributed systems
* decentralized networking
* protocol engineering
* concurrent runtime design
* cryptographic communication systems
* peer-to-peer architectures
* resilient communication ecosystems

---

# Current Priorities

Immediate focus:

1. Persistent peer loading and saving
2. Session deduplication
3. Automatic reconnect system
4. Peer introduction system
5. Local encryption vault
6. Encryption integration
7. Runtime stability improvements
8. Group communication support

---

# Current Milestones

Completed:

✓ IPv6 networking

✓ Multi-peer communication

✓ Heartbeat presence system

✓ Identity persistence

✓ Identity handshake

✓ Direct messaging

✓ Chat persistence

✓ Conversation history retrieval

Next Milestone:

→ Persistent peer graph (`peers.json`)
