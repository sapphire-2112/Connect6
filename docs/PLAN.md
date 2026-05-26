# Connect6 Development Plan

# Vision

Connect6 is a privacy-first decentralized peer-to-peer communication platform focused on:

- user-owned identity
- decentralized trust relationships
- encrypted communication
- local-only encrypted storage
- minimal metadata exposure

The goal is to build a communication ecosystem that avoids centralized identity systems, centralized message storage, and large-scale metadata collection.

---

# Current Architecture

Current system supports:

- IPv6 TCP peer communication
- decentralized node runtime
- concurrent peer handling
- structured JSON protocol messaging
- manual peer connectivity
- Dockerized distributed testing

---

# Current Project Structure

```text
cmd/node/
network/
peer/
protocol/
```

---

# Current Features

- multi-peer node connections
- manual peer connection via IPv6
- structured protocol messages
- concurrent message receiving
- peer management
- command-driven CLI

---

# Current Commands

/connect [ipv6]:8080
/peers

normal text = chat message

---

# Immediate Next Goals

## Networking
- stabilize peer connection lifecycle
- peer disconnect handling
- peer cleanup
- better terminal UI

## Identity
- cryptographic identity generation
- persistent node identity
- public/private key pairs

## Protocol
- protocol versioning
- message IDs
- targeted peer messaging
- peer announcement packets

## Security
- encrypted transport
- authenticated peer handshakes
- local encrypted storage

---

# Long-Term Goals

- decentralized trust propagation
- invite-based onboarding
- encrypted local vault
- peer discovery through trusted peers
- NAT traversal
- relay routing
- resilient decentralized communication graph

---

# Current Phase

Phase 1:
Core decentralized node runtime and transport foundation.