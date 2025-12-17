# 09 — Observability + AIOps: reduce MTTR, not just add dashboards

**Goal:** shorten time-to-detect and time-to-recover by making failures explain themselves.

Observability is not “more dashboards.” It’s the ability to answer new questions under stress:
- “What broke?”
- “Who is affected?”
- “What changed?”
- “What’s the fastest safe mitigation?”

## What you’ll learn

- how logs/metrics/traces work together (correlation is the superpower)
- what “good telemetry” looks like for microservices
- where AIOps helps (and where it can mislead)

## The three pillars (and why correlation matters)

- **Logs**: high detail, high volume (best for debugging)
- **Metrics**: low detail, low volume (best for alerting and trends)
- **Traces**: causal chain across services (best for distributed latency and failure)

You don’t “pick one.” You correlate them with:
- trace IDs propagated end-to-end
- consistent service naming and environment labels
- resource attributes (region, cluster, version)

## What “good telemetry” looks like (a concrete checklist)

- every request has a trace ID
- errors include domain context (not stack traces only)
- key business events are measured (orders/min, checkout failures)
- dashboards match user journeys (not just CPU graphs)
- alerts have runbooks and point to a likely action

## Instrumentation principles (keep it consistent)

- Use `OpenTelemetry` conventions where possible.
- Prefer structured logs (JSON) with stable fields.
- Avoid high-cardinality metric labels (they will melt your backend).
- Sample intelligently: capture errors and slow traces more often than the happy path.

## AIOps (practical, not magical)

Use automation to:
- group alerts (deduplicate)
- detect anomalies (baseline + seasonality)
- propose likely root causes (top correlated signals)

But keep humans in the loop:
- AIOps suggests, SRE approves action

## Tracing patterns (Tempo/Jaeger/any trace store)

Tracing shines when you:
- sample intelligently (tail-based sampling for errors/slow requests)
- link logs-to-traces and traces-to-logs
- propagate baggage/correlation IDs for key business dimensions

## Hands-on (this repo)

- Request logging middleware example: [`reference-implementation/services/catalog-go/main.go`](../reference-implementation/services/catalog-go/main.go)
- Synthetic traffic generator (useful for validating dashboards): [`reference-implementation/tests/k6/smoke.js`](../reference-implementation/tests/k6/smoke.js)
- Checklists to operationalize telemetry: [`docs/99-checklists.md`](99-checklists.md)

## Exercises (30–60 min)

- Define RED metrics for one service endpoint (rate, errors, duration) and how you would alert on them.
- Add a trace/request ID strategy to your services (even if only in logs), and decide how clients propagate it.

---

**Prev:** [08 — Resilience + chaos](08-resilience-chaos.md) · **Next:** [10 — Advanced testing](10-testing.md)
