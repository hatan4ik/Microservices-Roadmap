# Checklists (print these)

## Service readiness checklist
- [ ] Clear bounded context + owner (team/on-call)
- [ ] Health endpoints (liveness/readiness) and meaningful startup behavior
- [ ] Timeouts on all outbound calls (and sane retry rules)
- [ ] Structured logs + request/trace IDs
- [ ] Metrics for RED (rate, errors, duration)
- [ ] SLO defined (even a draft) for at least one critical journey
- [ ] CI: unit tests + lint + dependency scan
- [ ] Deploy manifest includes resource requests/limits
- [ ] PDB + HPA where appropriate
- [ ] Runbook exists (top failure modes + mitigations)

## API readiness checklist
- [ ] Contract documented (`OpenAPI`/proto) with examples
- [ ] Backward/forward compatibility rules defined
- [ ] Idempotency strategy for “create” and retryable operations
- [ ] Error semantics standardized (codes, retry guidance, correlation IDs)
- [ ] Rate limiting / quota policy defined (even if “not yet enforced”)
- [ ] Deprecation process documented (timeline + communication)

## Event readiness checklist
- [ ] Event envelope includes: `id`, `type`, `version`, `time`, `trace_id`
- [ ] Schema evolution strategy (what’s allowed / what breaks)
- [ ] Idempotent consumers (dedup key + storage strategy)
- [ ] Replay strategy documented (how to rebuild projections safely)
- [ ] Dead letter / poison message policy (who triages, how often, what actions)

## Kubernetes readiness checklist
- [ ] Resource requests/limits defined (and reviewed)
- [ ] Readiness/liveness probes reflect real health
- [ ] Graceful shutdown implemented (termination handling)
- [ ] HPA signal chosen (CPU, latency, queue depth)
- [ ] PDB defined if availability matters
- [ ] NetworkPolicy baseline applied (least privilege)

## Security checklist (baseline)
- [ ] Authentication everywhere
- [ ] Authorization per action (not “service-level allow-all”)
- [ ] Secrets not in git; rotation process exists
- [ ] SBOM produced for builds
- [ ] Image scanning gates releases (with clear exception policy)
- [ ] NetworkPolicies enforce least privilege

## Chaos checklist
- [ ] Hypothesis written
- [ ] Blast radius defined
- [ ] Abort conditions defined
- [ ] Success metric defined
- [ ] Results written down + follow-up actions

---

If you want, convert these into GitHub issue templates.
