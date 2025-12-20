# Lab 03: Kubernetes Troubleshooting Challenge

**Goal:** Apply the "Systematic Troubleshooting Framework" to fix a broken application.

**Scenario:** The Payments team pushed a new microservice. They claim "it works on my machine", but it fails in the cluster. Your job is to make it accessible via `curl`.

## Prerequisites

- A running local cluster (kind/minikube) - see [Lab 02](02-k8s-local.md).
- `kubectl` configured.

## 1) Deploy the Broken App

Apply the manifest that contains intentional errors:

```bash
kubectl apply -f ../reference-implementation/k8s/troubleshooting/broken-app.yaml
```

## 2) Phase 1: The Startup Failure

Check the pod status:

```bash
kubectl -n troubleshooting get pods
```

You should see `CreateContainerConfigError` or `Pending`.

**Investigation:**
1.  Describe the pod: `kubectl -n troubleshooting describe pod <pod-name>`
2.  Look at the "Events" section.
3.  What is missing?

**Fix:**
Create the missing resource. (Hint: It expects a Secret).

```bash
kubectl -n troubleshooting create secret generic payment-secrets --from-literal=api_key=supersecret
```

Wait a moment. The pod status should change.

## 3) Phase 2: The Crash Loop

Now check the pod again:

```bash
kubectl -n troubleshooting get pods -w
```

It might go to `Running`, but soon it will likely hit `Restarts: 1`, then `CrashLoopBackOff`.

**Investigation:**
1.  Check logs: `kubectl -n troubleshooting logs <pod-name>` (It might be empty if nginx started fine but was killed).
2.  Check previous logs: `kubectl -n troubleshooting logs <pod-name> --previous` (Still might be fine).
3.  Describe the pod again: `kubectl -n troubleshooting describe pod <pod-name>`
4.  Look at "Events". Do you see `Unhealthy`?
5.  Check the `Liveness` probe configuration in the manifest vs the container port.

**Fix:**
Edit the deployment to fix the Liveness Probe port.

```bash
kubectl -n troubleshooting edit deploy payments
```

*Change `port: 8080` to `port: 80` in the livenessProbe section.*

Wait for the rollout to finish. The pod should stay `Running`.

## 4) Phase 3: The Network Black Hole

Now try to access the service.

First, port-forward the **Service** (not the pod, we want to test service discovery):

```bash
kubectl -n troubleshooting port-forward svc/payments-svc 8080:80
```

In another terminal:

```bash
curl localhost:8080
```

*Error: Empty reply from server / Connection refused / Hanging.*

**Investigation:**
1.  Check if the Service has endpoints:
    ```bash
    kubectl -n troubleshooting get ep payments-svc
    ```
2.  If the `ENDPOINTS` column is `<none>`, the Service isn't finding any pods.
3.  Compare the Service `selector` with the Pod `labels`.
    - Service selector: `kubectl -n troubleshooting get svc payments-svc -o yaml`
    - Pod labels: `kubectl -n troubleshooting get pod --show-labels`

**Fix:**
Edit the Service to match the Pod's label.

```bash
kubectl -n troubleshooting edit svc payments-svc
```

*Change `app: payment-service` to `app: payments`.*

## 5) Verification

Check endpoints again:

```bash
kubectl -n troubleshooting get ep payments-svc
```

Now `curl localhost:8080` (with port-forward running) should return the Nginx default page.

## 6) Cleanup

```bash
kubectl delete namespace troubleshooting
```

---

**Summary:** You just debugged the three most common K8s issues:
1.  **Config Error:** Missing dependencies (Secret/ConfigMap).
2.  **Runtime/Health Error:** Bad probes killing a healthy app.
3.  **Discovery Error:** Mismatched selectors.
