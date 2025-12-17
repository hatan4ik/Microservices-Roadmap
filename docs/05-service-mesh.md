# Service Mesh Basics: what you gain, what you pay

A service mesh is a **network middleware layer**: retries, timeouts, mTLS, traffic shifting, and telemetry.

## When a mesh is worth it
- you need mTLS everywhere with consistent identity
- you need progressive delivery (canary) at L7
- you need uniform telemetry without rewriting all apps

## Costs
- operational complexity
- sidecar/resource overhead
- new failure modes (mesh control plane)

## Start small
1. enforce timeouts + retries in code first
2. add mesh for mTLS + traffic shaping later

Examples:
- `reference-implementation/mesh/` (DestinationRule/VirtualService-style examples)
- `reference-implementation/k8s/networkpolicy.yaml` (zero-trust baseline)

Next: `docs/06-api-management.md`
