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
* Background lightweight peer sessions
* Heartbeat-based online/offline presence tracking
* Active conversation targeting using `/use`
* Local chat persistence
* Conversation history retrieval
* Modular command architecture
* Dockerized distributed network testing
* Concurrent message receiving using Go routines
* Structured JSON-based protocol messaging

---

# Current Runtime Commands

```bash
/connect [ipv6]:8080
```

Connect to a peer node.

```bash
/peers
```

View currently known peer sessions and online status.

```bash
/use [peer_id]
```

Switch active conversation target.

```bash
/disconnect [peer_id]
```

Disconnect active transport session from a peer.

```bash
/history [peer_id]
```

View locally stored conversation history with a peer.

---

# Current Architecture

Connect6 currently separates communication into multiple layers:

## Identity Layer

Provides persistent node identities independent of transport sessions.

Responsibilities:

* node identity generation
* identity persistence
* identity exchange during connection establishment

## Peer Relationship Layer

Represents known or trusted peers within the decentralized graph.

## Background Session Layer

Maintains lightweight persistent peer sessions for:

* heartbeat exchange
* online presence visibility
* decentralized connectivity continuity

## Active Conversation Layer

Allows targeted peer messaging without broadcasting messages to every connected node.

## Local Persistence Layer

Provides local-first storage for:

* node identities
* conversation history
* future peer metadata persistence

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

* Peer metadata is not yet persisted across restarts
* Duplicate peer sessions are possible
* No cryptographic identity verification yet
* No encrypted messaging yet
* No automatic peer reconnect
* No decentralized peer discovery
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

Current stage:

Identity-based decentralized communication runtime with persistent local chat storage.

Implemented:

* IPv6 networking
* multi-peer communication
* heartbeat presence tracking
* identity persistence
* identity handshake protocol
* targeted messaging
* local chat persistence
* conversation history retrieval

Current focus areas:

* peer persistence
* session deduplication
* peer introduction system
* automatic reconnect
* encrypted local storage
* decentralized graph evolution
* protocol foundation development
