# API Management: gateways, GraphQL gateways, and Zero Trust enforcement

## Gateways (the boring but essential layer)
A gateway is your “front door”:
- authn/authz
- rate limiting
- routing
- request/response transformations
- observability headers

## “AI-powered gateways” (how to think about it safely)
You can use ML/AI for:
- anomaly-based rate limits (adaptive)
- bot detection
- WAF tuning suggestions
- intelligent routing during brownouts

But keep the **policy decision deterministic**:
- AI suggests, humans approve, policies enforce

## GraphQL federation
Federation is great when:
- clients need one query surface
- teams own subgraphs (autonomy)
- you control schema evolution rules

Avoid:
- building a single “supergraph team” that becomes a bottleneck

## Zero trust enforcement
Assume no network is trusted:
- authenticate every call
- authorize every action
- encrypt in transit
- segment with NetworkPolicies (and mTLS when possible)

See:
- `reference-implementation/k8s/networkpolicy.yaml`
- `reference-implementation/policies/` for policy-as-code examples

Next: `docs/07-devsecops.md`
