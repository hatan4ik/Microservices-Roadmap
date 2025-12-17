# 07 — DevSecOps: make the secure path the easy path

**Goal:** reduce risk without slowing delivery by turning security requirements into fast, automated feedback.

## What you’ll learn

- what “minimum viable security automation” looks like for microservices
- how to handle secrets safely across environments
- how policy-as-code prevents drift and enforces standards

## CI/CD security (minimum baseline)

- **SAST** (language scanners) for obvious unsafe patterns
- **Dependency scanning** (SBOM + CVE) for known vulnerable libraries
- **Container image scanning** (base image + packages + app deps)
- **IaC scanning** (K8s manifests, Terraform, Helm)
- **Signing/attestation** (recommended): prove what you built and how

Practical rule: security checks should be **fast**, **actionable**, and **blocking only when justified**.

## Supply-chain hygiene (modern requirement)

- pin versions (don’t float to “latest”)
- minimize base images and attack surface
- produce an SBOM per build
- ensure provenance (who built it, from what source, with what steps)

## Secrets management (keep it boring)

Anti-pattern: plaintext secrets in git or in Kubernetes manifests.

Preferred progression:
1. **Dev**: `Secret` for local-only development (still avoid committing real secrets)
2. **Prod**: External Secrets (cloud secrets manager) + least-privilege IAM
3. **High maturity**: Vault + dynamic creds + short TTL + rotation

Remember: Kubernetes `Secret` is not “magically safe” by default (treat it as a delivery mechanism, not a vault).

## Policy as Code

Two good layers:
- **OPA/Gatekeeper**: flexible and expressive (Rego)
- **Kyverno**: Kubernetes-native, approachable for platform teams

Use policy for standards that should never rely on “remembering”:
- no `latest` tags
- require resource requests/limits
- enforce baseline labels/annotations

## Hands-on (this repo)

- CI pipeline example: [`reference-implementation/.github/workflows/ci.yaml`](../reference-implementation/.github/workflows/ci.yaml)
- OPA example policy: [`reference-implementation/policies/opa/deny-latest-tag.rego`](../reference-implementation/policies/opa/deny-latest-tag.rego)
- Kyverno example policy: [`reference-implementation/policies/kyverno/require-requests-limits.yaml`](../reference-implementation/policies/kyverno/require-requests-limits.yaml)
- Secrets guidance: [`reference-implementation/policies/secret-handling.md`](../reference-implementation/policies/secret-handling.md)

## Exercises (45–90 min)

- Add a policy you wish your org had (pick one):
  - deny privileged pods
  - require non-root containers
  - require readiness probes
- Define a secrets threat model for a single service:
  - what secrets exist
  - where they live per environment
  - how they rotate

---

**Prev:** [06 — API management](06-api-management.md) · **Next:** [08 — Resilience + chaos](08-resilience-chaos.md)
