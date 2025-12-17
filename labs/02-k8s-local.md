# Lab 02: Deploy to Kubernetes locally (kind or minikube)

**Goal:** run the reference services in a local Kubernetes cluster and validate them via `kubectl port-forward`.

This lab assumes you already ran (or can run) the local stack in Docker: `labs/01-local-run.md`.

## Prerequisites

- `kubectl`
- A local cluster:
  - **kind** (recommended for a quick local cluster), or
  - **minikube**
- Docker (for building images)

Optional:
- `grpcurl` (to call the gRPC service)

## 1) Create a local cluster

### Option A — kind

```bash
kind create cluster --name mszth
kubectl config use-context kind-mszth
```

### Option B — minikube

```bash
minikube start
```

## 2) Build container images

The Kubernetes manifests reference these images:
- `mszth/catalog:dev`
- `mszth/orders:dev`

Build them from repo root:

```bash
cd reference-implementation
make build
```

## 3) Load images into the cluster

### Option A — kind

```bash
kind load docker-image mszth/catalog:dev mszth/orders:dev --name mszth
```

### Option B — minikube

```bash
minikube image load mszth/catalog:dev
minikube image load mszth/orders:dev
```

If you see `ErrImagePull` or `ImagePullBackOff`, the cluster can’t see your local images. Re-run the load step above.

## 4) Deploy (minimal)

Deploy just the core manifests (namespace, services, deployments):

```bash
cd reference-implementation
make k8s-apply-minimal
kubectl -n mszth get pods
```

Wait for the deployments to become ready:

```bash
kubectl -n mszth rollout status deploy/catalog
kubectl -n mszth rollout status deploy/orders
```

## 5) Validate via port-forward

### Catalog (HTTP)

```bash
kubectl -n mszth port-forward svc/catalog 8080:80
```

In another terminal:

```bash
curl http://localhost:8080/healthz
curl http://localhost:8080/v1/products
```

### Orders (gRPC)

```bash
kubectl -n mszth port-forward svc/orders 50051:50051
```

Then (optional, using `grpcurl`):

```bash
grpcurl -plaintext -d '{"customer_id":"c1","item_ids":["p1","p2"]}' localhost:50051 orders.v1.OrdersService/CreateOrder
```

## 6) Cleanup

```bash
cd reference-implementation
make k8s-delete-minimal
```

## Optional: apply the “add-ons”

The `reference-implementation/k8s/` folder also includes examples like an HPA, a PDB, and a default-deny `NetworkPolicy`.

If you want to apply everything:

```bash
cd reference-implementation
make k8s-apply
```

Notes:
- HPA typically needs `metrics-server`; if it’s not installed, the HPA may show warnings but the services still run.
- Some clusters enforce `NetworkPolicy` strictly; a default-deny policy can make services unreachable until you add explicit allow policies (treat it as a hardening exercise).

