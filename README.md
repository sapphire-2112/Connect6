# Connect6

Connect6 is a privacy-first decentralized peer-to-peer communication runtime focused on direct node-to-node communication, local identity ownership, and persistent decentralized peer relationships without relying on centralized messaging infrastructure.

Unlike traditional messaging platforms that depend on centralized servers, cloud-stored conversations, and platform-controlled identities, Connect6 is designed around direct peer communication where nodes communicate through lightweight persistent sessions over IPv6 TCP networking.

The project explores decentralized communication architecture, distributed systems concepts, peer runtime management, and protocol engineering using Go.

---

# Current Features

* Direct IPv6 peer-to-peer TCP communication
* Multi-peer concurrent node runtime
* Persistent node identities
* Identity handshake protocol
* Persistent peer storage (`peers.json`)
* Automatic peer loading on startup
* Automatic peer reconnect attempts
* Heartbeat-based online/offline presence tracking
* Active conversation targeting using `/use`
* Local chat persistence
* Conversation history retrieval
* Peer discovery protocol (`/peersof`)
* Modular command architecture
* Dockerized distributed network testing
* Concurrent message receiving using Go routines
* Structured JSON-based protocol messaging

---

# Current Runtime Commands

### `/connect [ipv6]:8080`

Connect to a peer node.

### `/peers`

View currently known peers and online status.

### `/peersof [peer_id]`

Request and view the peers known by another connected node.

### `/use [peer_id]`

Switch active conversation target.

### `/disconnect [peer_id]`

Disconnect a transport session from a peer.

### `/history [peer_id]`

View locally stored conversation history with a peer.

---

# Current Architecture

Connect6 separates communication into multiple layers.

## Identity Layer

Provides persistent node identities independent of transport sessions.

Responsibilities:

* node identity generation
* identity persistence
* identity exchange during connection establishment

---

## Peer Relationship Layer

Represents known peers within the decentralized communication graph.

Responsibilities:

* peer persistence
* peer loading on startup
* peer discovery
* peer relationship tracking

---

## Background Session Layer

Maintains lightweight persistent peer sessions for:

* heartbeat exchange
* online presence visibility
* automatic reconnect attempts
* decentralized connectivity continuity

---

## Active Conversation Layer

Allows targeted peer messaging without broadcasting messages to every connected node.

---

## Local Persistence Layer

Provides local-first storage for:

* node identities
* peer metadata
* conversation history

---

# Technologies Used

* Go
* TCP Networking
* IPv6
* Goroutines
* JSON Protocol Messaging
* Docker
* Concurrent Runtime Design

---

# Current Limitations

The project is still in active development.

Current limitations include:

* Duplicate peer sessions are possible
* No cryptographic identity verification yet
* No encrypted messaging yet
* No trust management system
* No connection request workflow
* No peer acceptance/rejection mechanism
* No NAT traversal support
* No offline message synchronization
* No group communication support

---

# Long-Term Vision

Connect6 aims to evolve into a decentralized communication ecosystem focused on:

* user-owned identities
* cryptographic identity verification
* encrypted local communication vaults
* decentralized trust propagation
* direct peer communication
* metadata minimization
* local-first encrypted storage
* decentralized peer introduction systems
* resilient distributed communication architecture

The long-term goal is to provide a privacy-centric alternative to centralized messaging systems while serving as a deep exploration into distributed systems, protocol engineering, decentralized networking, and secure communication design.

---

# Project Status

### Current Stage

Identity-based decentralized communication runtime with persistent peer relationships and peer discovery capabilities.

### Implemented

* IPv6 networking
* Multi-peer communication
* Heartbeat presence tracking
* Identity persistence
* Identity handshake protocol
* Persistent peer storage
* Automatic peer loading
* Automatic reconnect
* Targeted messaging
* Local chat persistence
* Conversation history retrieval
* Peer discovery protocol (`/peersof`)

### Current Focus Areas

* Peer introduction system
* Connection request workflow
* Peer acceptance/rejection system
* Session deduplication
* Cryptographic identity verification
* Message signing
* Encrypted local storage
* Trust graph evolution
* Group communication support
