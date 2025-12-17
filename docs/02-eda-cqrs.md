# Event-Driven Architecture (EDA) and CQRS without the hype

EDA and CQRS are power tools. Use them when you need the capabilities they provide.

## When EDA shines
- services must evolve independently
- you need asynchronous integration
- you need auditability (event log)
- you want fan-out to multiple consumers

## When EDA hurts
- you need strict, immediate consistency everywhere
- your org can’t operate asynchronous systems yet
- you don’t have tracing/observability maturity

## CQRS: separate reads from writes
**Command side**: validates invariants, writes the source of truth  
**Query side**: optimized projections for reads (often denormalized)

### The real reason CQRS exists
Not “microservices.” It exists because:
- read models and write models have *different shapes*
- read traffic can dwarf write traffic
- you want independent scaling

## Exactly-once myths
Most systems should design for:
- **at-least-once delivery**
- **idempotent consumers**
- **deduplication keys**
- **replayability**

## Reference patterns
- Outbox + relay (publish events after DB commit)
- Saga (orchestration vs choreography)
- Idempotency keys for commands

See examples:
- `reference-implementation/services/catalog/` (outbox concept notes)
- `reference-implementation/infra/` (Kafka-compatible broker)

Next: `docs/03-communication.md`
