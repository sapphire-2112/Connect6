# Connect6 Development Plan

# Current Stage

Connect6 currently functions as a decentralized multi-peer communication runtime where independently operating nodes communicate directly over IPv6 TCP connections.

Implemented capabilities include:

* concurrent peer handling
* active peer sessions
* heartbeat presence tracking
* targeted messaging
* decentralized node runtime behavior
* Docker-based distributed testing

---

# Core Architectural Goals

## Decentralized Identity

Move from transport-based identity (`IP:PORT`) toward persistent cryptographic node identities.

Planned:

* public/private key identities
* cryptographic peer verification
* identity persistence across network changes

---

# Persistent Peer Graph

Build a decentralized trust-based communication graph where:

* peers persist locally
* trusted nodes survive restarts
* relationships are user-controlled

Planned:

* automatic peer loading
* peer metadata persistence
* trust management

---

# Secure Communication Layer

Introduce encrypted peer-to-peer communication.

Planned:

* end-to-end encryption
* encrypted local vault
* encrypted protocol payloads
* secure session establishment

---

# Presence & Session Management

Expand runtime communication behavior.

Planned:

* automatic background reconnect
* session deduplication
* decentralized presence propagation
* heartbeat optimization

---

# Messaging System Evolution

Improve messaging architecture.

Planned:

* local chat persistence
* conversation history
* multi-conversation handling
* offline message handling
* peer conversation contexts

---

# Decentralized Networking Goals

Long-term networking direction includes:

* decentralized peer discovery
* invite-based onboarding
* peer introductions
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

1. Stable peer identity system
2. Session deduplication
3. Persistent peer loading
4. Automatic reconnect system
5. Local message persistence
6. Encryption integration
7. Runtime stability improvements
