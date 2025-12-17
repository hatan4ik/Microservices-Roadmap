# Secrets handling guidance

**Goal:** developers move fast without leaking secrets.

## Golden rules
- secrets never live in git
- secrets are short-lived when possible
- production access is audited
- apps read secrets the same way in every environment

## Preferred patterns
1) Dev: Kubernetes Secret (non-sensitive, local only)
2) Prod: External Secrets (cloud secrets manager)
3) High maturity: Vault dynamic creds + TTL + rotation

## Examples
- Kyverno policy: block plaintext secrets (see `kyverno/`)
- OPA policy: require labels/annotations for compliance (see `opa/`)
