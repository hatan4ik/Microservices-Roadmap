# Reference Implementation Map

This repo includes a small “reference implementation” that you can run locally and use as a sandbox while reading the chapters.

**Start here:** [Lab 01 — Run the reference system locally](../labs/01-local-run.md)

If you want Kubernetes next: [Lab 02 — Deploy to Kubernetes locally](../labs/02-k8s-local.md)

## What’s inside `reference-implementation/`

- **Services (Go)**
  - Catalog (HTTP): [`reference-implementation/services/catalog-go/`](../reference-implementation/services/catalog-go/)
  - Orders (gRPC): [`reference-implementation/services/orders-go/`](../reference-implementation/services/orders-go/)
- **Local infrastructure (Docker Compose)**: [`reference-implementation/infra/docker-compose.yml`](../reference-implementation/infra/docker-compose.yml)
  - Postgres + Redpanda (Kafka-compatible broker)
- **Kubernetes manifests**: [`reference-implementation/k8s/`](../reference-implementation/k8s/)
  - deployments, services, HPA/PDB, and a baseline `NetworkPolicy`
- **GitOps example**: [`reference-implementation/gitops/`](../reference-implementation/gitops/)
- **Policy-as-code examples**: [`reference-implementation/policies/`](../reference-implementation/policies/)
  - OPA/Rego and Kyverno
- **Service mesh examples**: [`reference-implementation/mesh/`](../reference-implementation/mesh/)
- **Chaos experiments**: [`reference-implementation/chaos/`](../reference-implementation/chaos/)
- **Testing examples**: [`reference-implementation/tests/`](../reference-implementation/tests/)
  - Pact-style contract shape + k6 smoke test
- **Protobuf IDL**: [`reference-implementation/proto/`](../reference-implementation/proto/)

## How the docs map to the repo

### Labs (hands-on)

- Lab 01 → run locally (Docker Compose): [`labs/01-local-run.md`](../labs/01-local-run.md)
- Lab 02 → run on Kubernetes (kind/minikube): [`labs/02-k8s-local.md`](../labs/02-k8s-local.md)

### Chapters (concepts → artifacts)

- **01 — DDD** → service boundaries: [`reference-implementation/services/`](../reference-implementation/services/)
- **02 — EDA + CQRS** → local broker + projections as exercises: [`reference-implementation/infra/docker-compose.yml`](../reference-implementation/infra/docker-compose.yml)
- **03 — Communication** → gRPC contract + HTTP endpoint: [`reference-implementation/proto/orders.proto`](../reference-implementation/proto/orders.proto), [`reference-implementation/services/`](../reference-implementation/services/)
- **04 — Kubernetes + GitOps** → manifests + Argo CD app: [`reference-implementation/k8s/`](../reference-implementation/k8s/), [`reference-implementation/gitops/`](../reference-implementation/gitops/)
- **05 — Service mesh** → mesh configs + network policy baseline: [`reference-implementation/mesh/`](../reference-implementation/mesh/), [`reference-implementation/k8s/networkpolicy.yaml`](../reference-implementation/k8s/networkpolicy.yaml)
- **06 — API management** → zero-trust baseline + policy examples: [`reference-implementation/k8s/networkpolicy.yaml`](../reference-implementation/k8s/networkpolicy.yaml), [`reference-implementation/policies/`](../reference-implementation/policies/)
- **07 — DevSecOps** → CI + policies + secret handling: [`reference-implementation/.github/workflows/ci.yaml`](../reference-implementation/.github/workflows/ci.yaml), [`reference-implementation/policies/`](../reference-implementation/policies/)
- **08 — Resilience + chaos** → Chaos Mesh examples: [`reference-implementation/chaos/`](../reference-implementation/chaos/)
- **09 — Observability + AIOps** → validate signals using synthetic traffic: [`reference-implementation/tests/k6/smoke.js`](../reference-implementation/tests/k6/smoke.js)
- **10 — Testing** → Pact-style contracts + k6: [`reference-implementation/tests/`](../reference-implementation/tests/)

---

**Back:** [Docs index](README.md)
