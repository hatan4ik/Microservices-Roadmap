# Documentation Index

This folder is the core learning path for the repo. Each chapter is short, opinionated, and designed to be read in order.

## Suggested reading order

1. [00 — Roadmap & maturity model](00-roadmap.md)
2. [01 — Domain-Driven Design (DDD)](01-ddd.md)
3. [02 — Event-Driven Architecture (EDA) + CQRS](02-eda-cqrs.md)
4. [03 — Service communication](03-communication.md)
5. [04 — Kubernetes + GitOps](04-kubernetes-gitops.md)
6. [05 — Service mesh](05-service-mesh.md)
7. [06 — API management](06-api-management.md)
8. [07 — DevSecOps](07-devsecops.md)
9. [08 — Resilience + chaos engineering](08-resilience-chaos.md)
10. [09 — Observability + AIOps](09-observability-aiops.md)
11. [10 — Advanced testing](10-testing.md)

## Reference implementation

- Map (how chapters connect to code/manifests): [Reference implementation map](97-reference-implementation.md)
- Root folder: [`reference-implementation/`](../reference-implementation/)

## Supporting docs

- [Glossary](98-glossary.md)
- [Checklists](99-checklists.md)

## Labs

- Labs live in `labs/` (start here): [Labs index](../labs/README.md)
- [Lab 01 — Run the reference system locally](../labs/01-local-run.md)
- [Lab 02 — Deploy to Kubernetes locally](../labs/02-k8s-local.md)

## Contributing to docs

If you add or edit chapters, keep them:
- practical (decisions, trade-offs, pitfalls)
- tied to the repo (point to files/commands where possible)
- vendor-neutral by default (call out vendor specifics explicitly)

Use the chapter template: [`docs/_template.md`](./_template.md)
