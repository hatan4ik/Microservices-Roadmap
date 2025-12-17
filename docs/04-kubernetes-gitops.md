# Cloud-Native Infrastructure: Kubernetes, Operators, and GitOps

## Kubernetes isn’t the goal
Kubernetes is a **standard runtime**. The goal is:
- fast, safe delivery
- reliable operations
- repeatable environments

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
- enables progressive delivery

See:
- Argo CD application example: `reference-implementation/gitops/argocd-application.yaml`

## Multi-cloud / hybrid (the practical take)
Avoid “abstract everything” platforms.
Instead:
- standardize on **Kubernetes primitives**
- keep cloud-specific parts behind interfaces (storage class, IAM, DNS)

## Observability foundations
Before you scale:
- metrics (RED/USE)
- logs with trace IDs
- tracing across services
- SLOs per critical user journey

Next: `docs/05-service-mesh.md`
