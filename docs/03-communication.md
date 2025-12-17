# 03 — Service Communication: sync, async, streaming, and realtime

**Goal:** choose the right communication style per interaction, and design contracts that can evolve safely.

You don’t “pick one.” Healthy systems use **a portfolio** of communication styles with clear defaults.

## What you’ll learn

- how to decide between RPC, events, and streaming
- how contracts evolve without breaking consumers
- how to avoid “distributed monolith” failure modes

## The decision table (rule of thumb)

- **HTTP/REST**: human debuggable, broad tooling, good for public APIs
- **gRPC**: high throughput, strict contracts, great for internal RPC
- **Async events (Kafka/Pulsar)**: decoupling + fan-out + replay (integration)
- **Streaming RPC**: sustained data flows (telemetry, feeds, realtime backends)
- **GraphQL federation**: unify client query layer; keep subgraphs owned by teams
- **WebSockets/SSE**: realtime updates to clients
- **MQTT**: lightweight pub/sub for IoT and constrained devices

## Contract strategy (how you avoid “Friday night outages”)

- **External APIs**: versioned, backward compatible, documented (`OpenAPI`)
- **Internal RPC**: protobuf schemas + lint + breaking change checks
- **Events**: versioned envelope + schema evolution rules (registry if possible)

Practical rule: **compatibility is a process, not a file format**.

## Defaults that prevent most incidents

- Always set **timeouts** on outbound calls (client-side) and enforce **deadlines** (server-side).
- Retry only when requests are **idempotent** (or you have idempotency keys).
- Prefer **bulkheads** (limits) over “unbounded concurrency”.
- Keep user journeys shallow: avoid long chains of synchronous hops.

## Common traps (and mitigations)

- **Chatty microservices** (too many sync hops) → collapse the boundary or introduce async integration.
- **“Events as RPC”** (request/response over Kafka) → use RPC or a workflow engine unless you truly need async.
- **GraphQL becomes a new monolith** → federation works only when subgraphs remain independently owned.
- **Inconsistent error semantics** → standardize error codes, retry rules, and tracing headers.

## Hands-on (this repo)

- gRPC proto (internal contract example): [`reference-implementation/proto/orders.proto`](../reference-implementation/proto/orders.proto)
- HTTP endpoint example: [`reference-implementation/services/catalog-go/main.go`](../reference-implementation/services/catalog-go/main.go)
- Synthetic smoke check example: [`reference-implementation/tests/k6/smoke.js`](../reference-implementation/tests/k6/smoke.js)

## Exercises (30–60 min)

- For one domain interaction, decide: HTTP vs gRPC vs event. Write down latency, coupling, and failure-mode reasons.
- Draft a compatibility policy:
  - how long you support old versions
  - how you deprecate fields/endpoints
  - how you detect breaking changes in CI

---

**Prev:** [02 — EDA + CQRS](02-eda-cqrs.md) · **Next:** [04 — Kubernetes + GitOps](04-kubernetes-gitops.md)
