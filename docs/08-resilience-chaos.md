# 08 — Resilience + Chaos Engineering: prove your system can survive reality

**Goal:** build systems that fail predictably and recover quickly, then validate that with hypothesis-driven experiments.

## What you’ll learn

- the resilience primitives that prevent most cascading failures
- how to think in “error budgets” and user journeys
- how to run chaos experiments safely (without “YOLO outages”)

## The resilience toolkit (defaults)

- **Timeouts (always)**: every outbound call; budget per user journey
- **Retries (carefully)**: only for idempotent operations, with backoff and caps
- **Circuit breakers**: stop hammering a dependency that is failing
- **Bulkheads**: isolate resources (threads, pools, queues) to prevent cross-contamination
- **Rate limiting**: protect systems from overload and abuse
- **Load shedding**: degrade non-critical work to preserve core journeys

If you only implement one thing: **timeouts**. Most incidents are “slow dependency → thread exhaustion → outage”.

## A practical method: protect critical user journeys

1. Identify the top 3 journeys (e.g., browse catalog, create order, check order status).
2. Define an SLO per journey (availability + latency).
3. Allocate timeout budgets and concurrency limits that match the SLO.
4. Add graceful degradation: “what do we return when dependency X is down?”

## Chaos engineering (the point)

Chaos is not random failure. It’s **hypothesis-driven testing**:
- “If Pod X dies, the user journey still succeeds with <1% error.”

Good chaos experiments have:
- a hypothesis
- a blast radius
- abort conditions
- a measurable success metric
- a write-up with follow-ups

## Kubernetes chaos (example)

Chaos Mesh can:
- kill pods
- add latency
- drop packets
- stress CPU/memory

## Hands-on (this repo)

- Chaos experiments:
  - latency injection: [`reference-implementation/chaos/latency.yaml`](../reference-implementation/chaos/latency.yaml)
  - pod kill: [`reference-implementation/chaos/pod-kill.yaml`](../reference-implementation/chaos/pod-kill.yaml)
- Checklists to operationalize this: [`docs/99-checklists.md`](99-checklists.md)

## Exercises (45–90 min)

- Write a resilience plan for one dependency call:
  - timeout budget
  - retry policy (or none)
  - circuit breaker thresholds
  - fallback behavior
- Draft a chaos hypothesis and success metric for one of the provided experiments.

---

**Prev:** [07 — DevSecOps](07-devsecops.md) · **Next:** [09 — Observability + AIOps](09-observability-aiops.md)
