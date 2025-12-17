# Resilience and Chaos Engineering: prove your system can survive reality

## The resilience toolkit
- timeouts (always)
- retries with budgets (careful)
- circuit breakers
- bulkheads (resource isolation)
- rate limiting
- load shedding

## Chaos engineering (the point)
Chaos is not random failure. It’s **hypothesis-driven testing**:
- “If Pod X dies, the user journey still succeeds with <1% error.”

## Kubernetes chaos (example)
Chaos Mesh can:
- kill pods
- add latency
- drop packets
- stress CPU/memory

See:
- `reference-implementation/chaos/latency.yaml`
- `reference-implementation/chaos/pod-kill.yaml`

## Adaptive rate limiting and circuit breakers
- keep user-facing endpoints responsive
- degrade non-critical features first
- preserve core transactions

Next: `docs/09-observability-aiops.md`
