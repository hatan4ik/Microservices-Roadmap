# DevSecOps and Security Automation: make the secure path the easy path

## CI/CD security (minimum baseline)
- SAST (language scanners)
- dependency scanning (SBOM + CVE)
- container image scanning
- IaC scanning (K8s manifests, Terraform)
- signed artifacts (optional but recommended)

See:
- GitHub Actions workflow: `reference-implementation/.github/workflows/ci.yaml`

## Secrets management
Anti-pattern: putting secrets in env vars in plain YAML.

Preferred progression:
1. use K8s secrets for dev only
2. use external secrets (cloud secrets manager) for prod
3. for high maturity: Vault + dynamic creds + short TTL

Examples:
- `reference-implementation/policies/secret-handling.md`

## Policy as Code
Two good layers:
- **OPA/Gatekeeper**: flexible, expressive (Rego)
- **Kyverno**: Kubernetes-native, easy for platform teams

Examples:
- `reference-implementation/policies/opa/`
- `reference-implementation/policies/kyverno/`

Next: `docs/08-resilience-chaos.md`
