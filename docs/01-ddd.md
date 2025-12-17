# 01 — Domain-Driven Design (DDD) for Microservices

**Goal:** use DDD as a *boundary tool* so services map to business capabilities, stable ownership, and independent delivery.

DDD is not “use aggregates everywhere.” In microservices, DDD is primarily about **choosing boundaries that reduce coordination**.

## What you’ll learn

- how to find bounded contexts (and spot “fake services”)
- how data ownership and team ownership reinforce each other
- how to design integration points (APIs/events) without sharing internals

## Core idea: bounded contexts become system boundaries

A microservice should align to a **bounded context**: a domain area with its own language, rules, and data model.

In practice that boundary becomes:
- the **team boundary** (ownership and on-call)
- the **data boundary** (write ownership; read via contracts)
- the **deploy boundary** (independent release cadence)

### Signals your boundary is wrong

- shared DB tables across services
- APIs that expose internal tables (“CRUD services”)
- “simple change” requires 3+ teams every sprint
- you have to deploy 5 services together to ship one feature

## A practical slicing workshop (repeatable)

1. Write the domain’s top 20 business capabilities (verbs): *price*, *reserve inventory*, *authorize payment*, *ship order*.
2. Group them into 4–8 “capability clusters”.
3. For each cluster, define:
   - **ubiquitous language** (terms the business uses)
   - **invariants** (rules that must always hold on writes)
   - **data ownership** (what this context writes/owns)
   - key **commands** and **events**
4. Validate the cut with two tests:
   - **change test**: most changes live in one context
   - **runtime test**: a user journey is not a chain of 10 RPC calls

## Data ownership (non-negotiable if you want autonomy)

**One service owns one set of data.** Other services read via:
- APIs (sync), or
- events + local projections (async / CQRS)

If two services must update the same rows to make progress, you don’t have service autonomy yet.

## Patterns you’ll use constantly

- **Strangler fig**: migrate one capability at a time behind a stable interface
- **Anti-corruption layer (ACL)**: translate legacy concepts into your domain model
- **Outbox**: publish events reliably as part of the same DB transaction (covered in the EDA chapter)

## Hands-on (this repo)

- Diagram (bounded contexts + data ownership): [`diagrams/ddd-bounded-contexts.mmd`](../diagrams/ddd-bounded-contexts.mmd)
- Example service boundaries:
  - HTTP: [`reference-implementation/services/catalog-go/main.go`](../reference-implementation/services/catalog-go/main.go)
  - gRPC: [`reference-implementation/services/orders-go/main.go`](../reference-implementation/services/orders-go/main.go)
- Service readiness checklist: [`docs/99-checklists.md`](99-checklists.md)

## Exercises (30–60 min)

- Pick a system you know. Identify 3 bounded contexts and name the *team* that would own each.
- For one context, write:
  - 5 commands (writes)
  - 5 events (facts that happened)
  - 3 invariants (rules enforced on writes)
- Find one integration point that should be **async** and one that should be **sync**, and justify the trade-off.

## Further reading (optional)

- Eric Evans — *Domain-Driven Design*
- Vaughn Vernon — *Implementing Domain-Driven Design*
- Skelton & Pais — *Team Topologies* (how org design shapes boundaries)

---

**Prev:** [00 — Roadmap & maturity model](00-roadmap.md) · **Next:** [02 — EDA + CQRS](02-eda-cqrs.md)
