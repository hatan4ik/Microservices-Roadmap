# Checklists (print these)

## Service readiness checklist
- [ ] clear bounded context + ownership
- [ ] health endpoints (liveness/readiness)
- [ ] timeouts on outbound calls
- [ ] structured logs + trace IDs
- [ ] metrics for RED (rate, errors, duration)
- [ ] SLO defined (even a draft)
- [ ] CI: unit tests + lint + scan
- [ ] deploy manifest includes resource requests/limits
- [ ] PDB + HPA where appropriate
- [ ] runbook exists

## Event readiness checklist
- [ ] event envelope includes: id, type, version, time, trace_id
- [ ] schema versioning strategy
- [ ] idempotent consumers
- [ ] replay strategy documented
- [ ] dead letter / poison message plan

## Security checklist (baseline)
- [ ] authentication everywhere
- [ ] authorization per action (not per service)
- [ ] secrets not in git; rotated
- [ ] SBOM produced
- [ ] image scanning gates releases
- [ ] NetworkPolicies enforce least privilege

## Chaos checklist
- [ ] hypothesis written
- [ ] blast radius defined
- [ ] abort conditions defined
- [ ] success metric defined
- [ ] results written down + follow-up actions

---

If you want, convert these into GitHub issue templates.
