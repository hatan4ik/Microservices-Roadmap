# 04 — Cloud-Native Infrastructure: Kubernetes + GitOps

**Goal:** run services on a standard, repeatable runtime and ship changes safely with auditability and easy rollbacks.

Kubernetes isn’t the goal. Kubernetes is a **standard runtime**; the goal is:
- fast, safe delivery
- reliable operations
- repeatable environments

## What you’ll learn

- which Kubernetes primitives matter most for microservices
- what GitOps changes about day-to-day operations
- how operators reduce “day-2” toil for stateful systems

## Kubernetes primitives (the 20% that gives 80% of the value)

- **Deployment**: rollout strategy, health probes, resource requests/limits
- **Service**: stable discovery and load balancing
- **HPA**: basic autoscaling (CPU and beyond)
- **PDB**: keep availability during voluntary disruptions
- **NetworkPolicy**: least-privilege networking (baseline for zero trust)

Rule: if you can’t describe your service’s **health, scaling trigger, and blast radius**, you’re not “ready for K8s” yet.

## Operators and CRDs (why they matter)

Operators encode operational knowledge as code:
- backups, failovers, scaling rules
- day-2 automation (patching, migrations)

If you run stateful systems (DBs, Kafka, Redis), you will likely rely on operators.

## GitOps: one mental model

Git is the source of truth. The cluster converges to Git.

### Benefits

- auditable changes
- repeatable environments
- easy rollbacks
- enables progressive delivery and policy enforcement

## Multi-cloud / hybrid (the practical take)

Avoid “abstract everything” platforms. Instead:
- standardize on **Kubernetes primitives**
- keep cloud-specific parts behind interfaces (storage classes, IAM, DNS)

## Observability foundations (don’t skip)

Before you scale:
- metrics (RED/USE)
- logs with trace IDs
- tracing across services
- SLOs per critical user journey

If you don’t know what “healthy” looks like, autoscaling just makes failures happen faster.

## Hands-on (this repo)

- Kubernetes manifests: [`reference-implementation/k8s/`](../reference-implementation/k8s/)
- GitOps example (Argo CD Application): [`reference-implementation/gitops/argocd-application.yaml`](../reference-implementation/gitops/argocd-application.yaml)
- CI pipeline example: [`reference-implementation/.github/workflows/ci.yaml`](../reference-implementation/.github/workflows/ci.yaml)
- Apply manifests (from `reference-implementation/`): `make k8s-apply`

## Exercises (45–90 min)

- Deploy to a local cluster (kind/minikube), then:
  - verify readiness/liveness endpoints
  - observe HPA and PDB behavior
  - reason about the `NetworkPolicy` blast radius
- Add one operational requirement to the manifests:
  - resource requests/limits
  - pod anti-affinity
  - a more realistic HPA signal

---

**Prev:** [03 — Service communication](03-communication.md) · **Next:** [05 — Service mesh](05-service-mesh.md)
