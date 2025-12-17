# Domain-Driven Design for Microservices: bounded contexts that actually work

DDD is not “use aggregates everywhere.” In microservices, DDD is a **service boundary tool**.

## Core idea
A microservice should align to a **bounded context**: a domain area with its own language, rules, and data model.
That boundary becomes:
- the team boundary
- the data ownership boundary
- the deploy boundary

### Signals your boundary is wrong
- shared DB tables across services
- APIs that expose internal tables (“CRUD services”)
- changes require 3+ teams every sprint

## Practical slicing method (workshop)
1. Write your domain’s top 20 business capabilities (verbs).
2. Group them into 4–8 “capability clusters.”
3. For each cluster, define:
   - **ubiquitous language**
   - **invariants** (rules that must always hold)
   - **data ownership**
   - key commands/events

## Data ownership rule
**One service owns one set of data.**
Other services read via:
- APIs (sync), or
- events + local projections (async / CQRS)

## Patterns you’ll use constantly
- **Strangler fig**: migrate one capability at a time
- **Anti-corruption layer**: isolate legacy concepts
- **Outbox**: publish events reliably with DB transactions

## Exercises
- Take any system you know. Identify 3 bounded contexts.
- For one boundary, list the “commands” and “events”.

Next: `docs/02-eda-cqrs.md`
