# 06 — API Management: gateways, federation, and zero trust

**Goal:** make APIs safe to expose and safe to change — without turning your gateway into a single point of failure (or a bottleneck team).

## What you’ll learn

- what belongs at the edge (and what should stay in services)
- how to version and deprecate APIs without breaking consumers
- how “zero trust” shows up in API and network policy

## Gateways (the boring but essential layer)

A gateway is your “front door”:
- authentication and authorization
- rate limiting and quotas
- routing and service discovery integration
- request/response transformations (sparingly)
- observability headers (trace propagation, request IDs)

Good gateway rule: **centralize cross-cutting policy, not business logic**.

## API lifecycle: change without outages

- Document contracts (`OpenAPI`/protos), publish examples, and define deprecation rules.
- Prefer additive changes (new fields/endpoints) over breaking ones.
- Define error semantics and retry guidance (so every client doesn’t guess).

## GraphQL federation (when it helps)

Federation is great when:
- clients need one query surface
- teams own subgraphs (autonomy)
- you control schema evolution rules and performance budgets

Avoid:
- building a single “supergraph team” that becomes a bottleneck
- hiding N+1 calls and latency problems behind “one GraphQL request”

## “AI-powered gateways” (how to think about it safely)

You can use ML/AI for:
- anomaly-based rate limits (adaptive suggestions)
- bot detection and abuse heuristics
- WAF tuning suggestions
- intelligent routing during brownouts

But keep the **policy decision deterministic**:
- AI suggests, humans approve, policies enforce

## Zero trust enforcement (the practical version)

Assume no network is trusted:
- authenticate every call
- authorize every action (not “service-to-service is trusted”)
- encrypt in transit
- segment with `NetworkPolicy` (and mTLS when possible)

## Hands-on (this repo)

- Network segmentation baseline: [`reference-implementation/k8s/networkpolicy.yaml`](../reference-implementation/k8s/networkpolicy.yaml)
- Policy-as-code examples: [`reference-implementation/policies/`](../reference-implementation/policies/)

## Exercises (30–60 min)

- Write an API change policy for your team:
  - versioning strategy
  - deprecation timeline
  - compatibility checks in CI
- Pick one endpoint and define:
  - authn/authz requirements
  - rate limit policy
  - audit logging requirements

---

**Prev:** [05 — Service mesh](05-service-mesh.md) · **Next:** [07 — DevSecOps](07-devsecops.md)
