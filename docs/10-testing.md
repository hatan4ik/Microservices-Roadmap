# Advanced Testing: contracts, failover drills, and synthetic monitoring

## Contract testing (Pact)
Contract tests prevent “it worked in staging” surprises by validating:
- request/response shapes
- required fields
- semantic expectations

See:
- `reference-implementation/tests/pact/`

## Failover & recovery tests
Schedule them, don’t promise you’ll “do it later.”
- database failover simulation
- zone outage simulation
- message broker restart/recovery

## Synthetic monitoring
Synthetic checks catch regressions early:
- run in CI for PRs (lightweight)
- run continuously in prod (critical journeys)

See:
- `reference-implementation/tests/k6/smoke.js`

## The testing pyramid for microservices
- unit tests: fastest, most of your tests
- integration tests: service + DB
- contract tests: between services
- end-to-end tests: minimal set, stable environments
- chaos/resilience tests: scheduled, hypothesis-driven

Next: `docs/99-checklists.md`
