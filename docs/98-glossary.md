# Glossary (quick)

Short definitions for common terms used across the chapters.

- **Anti-corruption layer (ACL)**: translation layer that prevents legacy models from leaking into your domain model.
- **At-least-once delivery**: messages may be delivered more than once; consumers must be idempotent.
- **Bounded context**: a domain boundary with its own language and model (often maps to a service/team boundary).
- **Bulkhead**: isolate resources so one failing dependency doesn’t take down everything (pools/queues/limits).
- **Canary release**: gradually shift traffic to a new version and roll back quickly if signals degrade.
- **Circuit breaker**: stop calling a dependency that is failing to prevent cascading failure.
- **CQRS**: separate command (write) and query (read) models, often with projections for reads.
- **Distributed monolith**: many services that must deploy and change together (all the cost, none of the benefit).
- **DLQ (dead letter queue)**: place for messages that cannot be processed after retries; requires a handling policy.
- **Event envelope**: standard metadata around an event (id, type, version, time, trace_id, data).
- **GitOps**: Git is the source of truth for desired state; the cluster reconciles to it continuously.
- **HPA**: Kubernetes Horizontal Pod Autoscaler (scales replicas based on signals like CPU).
- **Idempotency**: an operation can be safely retried without changing the result beyond the first success.
- **mTLS**: mutual TLS; both client and server authenticate each other and encrypt traffic.
- **Outbox**: publish events reliably after DB commit (transactional consistency between DB write and event).
- **PDB**: Kubernetes Pod Disruption Budget (limits voluntary disruptions to maintain availability).
- **Policy as code**: enforce standards via declarative policies (OPA/Kyverno) instead of human memory.
- **RED**: rate, errors, duration (request-centric metrics).
- **Saga**: pattern for coordinating multi-step workflows across services (orchestration vs choreography).
- **SLO**: service level objective (target reliability/latency for a user journey).
- **Strangler fig**: incrementally replace a system by routing one capability at a time to the new implementation.
- **USE**: utilization, saturation, errors (resource-centric metrics).
