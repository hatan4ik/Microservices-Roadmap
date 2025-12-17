# Microservices Zero-to-Hero (O’Reilly-style training repo)

This repo is a **self-contained training path** for building microservices that scale: architecture, comms, cloud-native infra, DevSecOps, resilience, observability, and advanced testing.

It’s written like a practical “field guide”: **concepts → trade-offs → patterns → hands-on labs → reference implementation**.

---

## How to use this repo (recommended learning path)

1. Read the roadmap: `docs/00-roadmap.md`
2. Work through chapters in order (`docs/01-*` … `docs/10-*`)
3. Run the reference implementation locally: `reference-implementation/infra/docker-compose.yml`
4. Deploy to Kubernetes (kind/minikube or real cluster): `reference-implementation/k8s/`
5. Add governance, mesh, chaos, and CI: `reference-implementation/{policies,mesh,chaos,.github}`

---

## Repo structure

- `docs/` – chapters (O’Reilly-style narrative + checklists)
- `diagrams/` – Mermaid diagrams you can paste into GitHub or Docs
- `reference-implementation/`
  - `services/` – small services (Go HTTP + Go gRPC)
  - `proto/` – protobuf IDL
  - `infra/` – local dev stack (Postgres + Kafka-compatible broker + observability)
  - `k8s/` – Kubernetes manifests (deploy, svc, hpa, pdb, networkpolicy)
  - `gitops/` – Argo CD example
  - `policies/` – OPA/Rego + Kyverno examples
  - `mesh/` – Istio/Linkerd-style examples (traffic, mTLS)
  - `chaos/` – Chaos Mesh experiments
  - `tests/` – Pact-style contracts + k6 synthetic + failure drills
  - `.github/workflows/` – CI pipeline (build/test/scan)

---

## Prerequisites

- Docker + Docker Compose
- Go 1.22+ (for the reference services)
- kubectl + a cluster (kind/minikube is fine) for the K8s chapters

Optional (for deeper chapters):
- Argo CD or Flux
- Istio or Linkerd
- Chaos Mesh
- Prometheus + Grafana + Tempo/Loki

---

## What “good” looks like

By the end, you should be able to:
- slice services using **DDD bounded contexts**
- choose **sync vs async** communication and model contracts
- implement **EDA + CQRS** (and know when *not* to)
- ship safely with **GitOps**, **policy-as-code**, and **progressive delivery**
- keep systems stable using **SLOs**, **tracing**, and **chaos experiments**
- test at multiple layers: **unit → contract → integration → resilience**

---

## License

MIT (see `LICENSE`).
