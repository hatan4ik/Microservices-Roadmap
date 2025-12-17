# Service Communication: sync, async, streaming, and realtime

You don’t “pick one.” Most systems use **a portfolio** of communication styles.

## The decision table (rule of thumb)

- **HTTP/REST**: human debuggable, broad tooling, good for public APIs
- **gRPC**: high throughput, strict contracts, good for internal RPC
- **Async events (Kafka/Pulsar)**: decoupling + fan-out + replay
- **GraphQL federation**: unify query layer for clients; keep services autonomous
- **WebSockets**: client realtime updates
- **MQTT**: lightweight pub/sub for IoT and constrained devices

## Contract strategy
- **External APIs**: versioned, backward compatible, documented (OpenAPI)
- **Internal RPC**: protobuf schemas + lint + breaking change checks
- **Events**: schema registry (or at least versioned envelopes)

## Common traps
- “Chatty microservices”: too many sync hops
- “Events as RPC”: using Kafka for request/response (avoid unless needed)
- “GraphQL becomes a new monolith”: federation must keep subgraphs owned

Hands-on:
- gRPC proto: `reference-implementation/proto/orders.proto`
- WebSocket pattern notes: `docs/03-communication.md` (this file)

Next: `docs/04-kubernetes-gitops.md`
