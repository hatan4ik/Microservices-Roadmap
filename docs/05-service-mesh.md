# 05 — Service Mesh Basics: what you gain, what you pay

**Goal:** understand when a mesh meaningfully reduces risk/cost, and when it’s just complexity.

A service mesh is a **network middleware layer**: retries, timeouts, mTLS, traffic shifting, and telemetry — typically implemented via sidecars or “ambient” dataplanes.

## What you’ll learn

- what a mesh standardizes (and what still belongs in application code)
- why mTLS + identity is often the killer feature
- the operational costs and failure modes you must plan for

## When a mesh is worth it

- you need mTLS everywhere with consistent identity and authorization policy
- you need progressive delivery (canary) at L7 without every team re-implementing it
- you need uniform telemetry without rebuilding every client library

## When you should wait

- you don’t have timeouts/retries/circuit breakers in place at all
- the platform team can’t operate the control plane (upgrades, incidents, policy changes)
- service ownership is unclear (a mesh can’t fix unclear boundaries)

## Costs (be honest about them)

- operational complexity (control plane upgrades, policy debugging)
- sidecar/resource overhead (CPU/memory, latency)
- new failure modes (dataplane config, cert rotation, control plane outages)

## A safe adoption sequence

1. Enforce **timeouts** and **basic retry policies** in code first.
2. Add a mesh for **mTLS + identity** and **traffic shaping**.
3. Add policy enforcement (authz), then advanced routing (fault injection, mirroring).

## Hands-on (this repo)

- Mesh configuration examples: [`reference-implementation/mesh/`](../reference-implementation/mesh/)
- Baseline segmentation (zero trust starting point): [`reference-implementation/k8s/networkpolicy.yaml`](../reference-implementation/k8s/networkpolicy.yaml)

## Exercises (30–60 min)

- Choose one outbound call in your system. Define:
  - timeout budget
  - retry policy (or explicitly “no retries”)
  - circuit-breaker or bulkhead limit
- If you have Istio/Linkerd available, implement a simple canary route using the example manifests.

---

**Prev:** [04 — Kubernetes + GitOps](04-kubernetes-gitops.md) · **Next:** [06 — API management](06-api-management.md)
