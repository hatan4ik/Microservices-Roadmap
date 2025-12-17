# Microservices Zero-to-Hero

This repo is a **self-contained training path** for building microservices that scale: architecture, comms, cloud-native infra, DevSecOps, resilience, observability, and advanced testing.

It’s written like a practical “field guide”: **concepts → trade-offs → patterns → hands-on labs → reference implementation**.

---

## How to use this repo (recommended learning path)

1. Start at the docs index: [`docs/README.md`](docs/README.md)
2. Read the roadmap: [`docs/00-roadmap.md`](docs/00-roadmap.md)
3. Work through chapters in order (`docs/01-*` … `docs/10-*`)
4. Run the reference implementation locally: [`labs/01-local-run.md`](labs/01-local-run.md)
5. Deploy to Kubernetes (kind/minikube or real cluster): [`reference-implementation/k8s/`](reference-implementation/k8s/)
6. Add governance examples:
   - policies: [`reference-implementation/policies/`](reference-implementation/policies/)
   - mesh: [`reference-implementation/mesh/`](reference-implementation/mesh/)
   - chaos: [`reference-implementation/chaos/`](reference-implementation/chaos/)
   - CI: [`reference-implementation/.github/workflows/`](reference-implementation/.github/workflows/)

---

## Quickstart (local)

```bash
cd reference-implementation
make up
curl http://localhost:8080/healthz
curl http://localhost:8080/v1/products
make down
```

No `make`? Use Docker Compose directly:

```bash
docker compose -f reference-implementation/infra/docker-compose.yml up -d --build
docker compose -f reference-implementation/infra/docker-compose.yml down -v
```

## Repo structure

- `docs/` – chapters (O’Reilly-style narrative + checklists)
- `diagrams/` – Mermaid diagrams you can paste into GitHub or Docs
- `labs/` – hands-on walkthroughs (start with `labs/01-local-run.md`)
- `reference-implementation/`
  - `services/` – small services (Go HTTP + Go gRPC)
  - `proto/` – protobuf IDL
  - `infra/` – local dev stack (Postgres + Kafka-compatible broker)
  - `k8s/` – Kubernetes manifests (deploy, svc, hpa, pdb, networkpolicy)
  - `gitops/` – Argo CD example
  - `policies/` – OPA/Rego + Kyverno examples
  - `mesh/` – Istio/Linkerd-style examples (traffic, mTLS)
  - `chaos/` – Chaos Mesh experiments
  - `tests/` – Pact-style contracts + k6 synthetic + failure drills
  - `.github/workflows/` – CI pipeline (build/test/scan)

---

## Prerequisites

Minimum (local quickstart):
- Docker Desktop / Docker Engine with Docker Compose v2
- `make` (optional; you can use `docker compose` directly)

For running tests locally:
- Go 1.22+

For Kubernetes chapters:
- `kubectl` + a cluster (kind/minikube is fine)

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
