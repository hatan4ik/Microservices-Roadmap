# Secrets handling guidance

**Goal:** developers move fast without leaking secrets.

## Golden rules
- Secrets never live in git (including “temporary” test keys).
- Prefer short-lived secrets when possible (dynamic creds, rotated tokens).
- Production access is audited and least-privilege.
- Apps read secrets the same way in every environment (same interface, different backends).

## What counts as a secret?

Treat these as secrets:
- passwords, API keys, signing keys, encryption keys
- database connection strings (often include credentials)
- tokens/cookies that grant access

Treat these as configuration (still validate and protect, but not “secret”):
- feature flags (usually)
- non-sensitive URLs
- timeouts/limits

## Preferred patterns
1) Dev: Kubernetes Secret (non-sensitive, local only)
2) Prod: External Secrets (cloud secrets manager)
3) High maturity: Vault dynamic creds + TTL + rotation

## Operational checklist (minimum)

- rotate secrets (document who/when/how)
- define where secrets are injected (env vars vs mounted files vs sidecars)
- ensure logs never print secret material
- block secrets from being committed (pre-commit, CI scanning)

## Examples

This repo includes small, illustrative policy examples:
- Kyverno: require resource requests/limits (a common “secure-by-default” baseline): `kyverno/require-requests-limits.yaml`
- OPA/Rego: deny `:latest` image tags (supply chain hygiene): `opa/deny-latest-tag.rego`

If you want to extend this repo, good next policies include:
- deny plaintext `Secret` objects in non-dev namespaces
- require non-root containers
- require readiness probes and timeouts
