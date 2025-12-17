# 10 — Advanced Testing: contracts, drills, and synthetic monitoring

**Goal:** reduce production surprises by testing the things that actually break in distributed systems: contracts, dependencies, and failure modes.

## What you’ll learn

- where contract tests fit (and what they don’t replace)
- how to run failover/recovery drills as engineering work, not “a someday task”
- how synthetic monitoring closes the gap between CI and prod

## The testing pyramid for microservices (practical version)

- **Unit tests**: fastest; most of your tests live here
- **Integration tests**: service + DB/broker; validate real dependencies
- **Contract tests**: between services; stop breaking changes early
- **End-to-end tests**: minimal set; stable environments only
- **Resilience tests**: scheduled; hypothesis-driven (chaos/drills)

Rule: don’t use E2E tests to compensate for missing contracts and missing observability.

## Contract testing (Pact-style)

Contract tests prevent “it worked in staging” surprises by validating:
- request/response shapes
- required fields
- semantic expectations (e.g., “price must be positive”)

They are especially valuable when teams deploy independently.

## Failover & recovery tests (make them routine)

Schedule them; don’t promise you’ll “do it later”:
- database failover simulation
- zone outage simulation
- message broker restart/recovery

Write down results and add follow-up work, or the drill did not happen.

## Synthetic monitoring (CI + prod)

Synthetic checks catch regressions early:
- run in CI for PRs (lightweight)
- run continuously in prod (critical journeys)

## Hands-on (this repo)

- Pact-style contract example: [`reference-implementation/tests/pact/`](../reference-implementation/tests/pact/)
- k6 smoke test: [`reference-implementation/tests/k6/smoke.js`](../reference-implementation/tests/k6/smoke.js)
- Run Go unit tests (from `reference-implementation/`): `make test`

## Exercises (30–60 min)

- Write one consumer contract for an interaction you care about, then define how the provider verifies it in CI.
- Define one synthetic “golden journey” and decide:
  - what success looks like (latency + correctness)
  - what you alert on (and what you don’t)

---

**Prev:** [09 — Observability + AIOps](09-observability-aiops.md) · **Next:** [Checklists](99-checklists.md)
