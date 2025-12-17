# 02 — Event-Driven Architecture (EDA) + CQRS (without the hype)

**Goal:** use events and CQRS where they buy you autonomy, scalability, and auditability — without creating an un-debuggable “async spaghetti” system.

EDA and CQRS are power tools. If you don’t need the capabilities, prefer simpler sync APIs.

## What you’ll learn

- when async integration is worth the operational cost
- how to design events that survive schema evolution
- why “exactly-once” is rarely the right goal
- how CQRS projections work (and why they drift)

## When EDA shines

- services must evolve independently (looser coupling than RPC)
- you need asynchronous integration (don’t block user journeys on other teams)
- you need auditability / replay (event log as source of facts)
- you want fan-out to multiple consumers (search index, fraud, analytics)

## When EDA hurts (or should be delayed)

- you need strict, immediate consistency everywhere
- your org can’t operate async systems yet (on-call, dashboards, runbooks)
- you don’t have basic observability (trace IDs, consumer lag, DLQ visibility)

## CQRS: separate reads from writes

**Command side**: validates invariants, writes the source of truth  
**Query side**: optimized projections for reads (often denormalized)

### The real reason CQRS exists

Not “microservices.” CQRS exists because:
- read models and write models have *different shapes*
- read traffic can dwarf write traffic
- you want independent scaling (and different caching strategies)

## Event design: prefer “facts” over “requests”

Events should be facts that happened in a domain:
- ✅ `OrderPlaced`, `PaymentAuthorized`, `ShipmentDelivered`
- ❌ `PlaceOrderPlease`, `UpdateCustomerEmailNow`

If you find yourself “waiting for a response” via Kafka, you’re usually rebuilding RPC with worse ergonomics.

## Delivery semantics: the “exactly-once” trap

Most systems should design for:
- **at-least-once delivery**
- **idempotent consumers**
- **deduplication keys**
- **replayability**

Treat duplicates, retries, and out-of-order delivery as normal — because they are.

## Reference patterns (high signal)

- **Outbox + relay**: publish events after DB commit (transactional consistency)
- **Saga**: orchestration vs choreography for multi-step business processes
- **Idempotency keys**: make commands safe to retry
- **DLQ / poison message handling**: explicit policy for “cannot process”

## Hands-on (this repo)

- Diagram (EDA + CQRS projection flow): [`diagrams/eda-cqrs.mmd`](../diagrams/eda-cqrs.mmd)
- Kafka-compatible broker in the local stack: [`reference-implementation/infra/docker-compose.yml`](../reference-implementation/infra/docker-compose.yml) (Redpanda)
- Event readiness checklist: [`docs/99-checklists.md`](99-checklists.md)
- Starter services (patterns described in comments; exercises to implement):
  - [`reference-implementation/services/orders-go/main.go`](../reference-implementation/services/orders-go/main.go)
  - [`reference-implementation/services/catalog-go/main.go`](../reference-implementation/services/catalog-go/main.go)

## Exercises (45–90 min)

- Define an event envelope for this repo (include `id`, `type`, `version`, `time`, `trace_id`, `data`).
- Pick one domain flow (e.g., order creation). Write:
  - 3 events you would publish
  - 2 projections you would maintain
  - the idempotency/dedup strategy for each consumer
- Add a “failure mode table”: what happens when the consumer is down, slow, or errors repeatedly?

---

**Prev:** [01 — DDD](01-ddd.md) · **Next:** [03 — Service communication](03-communication.md)
