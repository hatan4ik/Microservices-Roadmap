# Observability and AIOps: reduce MTTR, not just add dashboards

## The three pillars (and why correlation matters)
- logs: high detail, high volume
- metrics: low detail, low volume
- traces: causal chain across services

You don’t “pick one.” You correlate them with:
- trace IDs
- consistent service naming
- resource attributes

## What “good telemetry” looks like
- every request has a trace ID
- errors include domain context (not stack traces only)
- key business events are measured (orders/min, checkout failures)

## AIOps (practical)
Use automation to:
- group alerts (deduplicate)
- detect anomalies (baseline + seasonality)
- propose likely root causes (top correlated signals)

But keep humans in the loop:
- AIOps suggests, SRE approves action

## Tempo / tracing patterns
Tempo (or any trace store) shines when you:
- sample intelligently (tail-based sampling for errors)
- link logs-to-traces and traces-to-logs

Next: `docs/10-testing.md`
