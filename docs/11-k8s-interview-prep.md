# 11 — Kubernetes Interview Prep: Deep Dive, On-Prem & Troubleshooting

**Goal:** Master the "below the watermark" details required for FAANG/MANGA systems design and troubleshooting interviews.

This chapter is less about "how to use K8s" and more about **"how it breaks and how to fix it"**, specifically in **on-prem/bare-metal** environments where you don't have a cloud provider managing the control plane.

## The Systematic Troubleshooting Framework

When an interviewer asks "A pod is down, how do you debug it?", do **not** just say "I check the logs". Use a systematic approach (bottom-up or top-down).

**The "5 Layers" Model:**

1.  **Cluster/Node Level:** Is the node ready? Is the API server responding? Is Etcd healthy?
2.  **Network Level:** Can nodes talk to each other? Is CNI happy? Is DNS resolving?
3.  **Controller Level:** Is the Deployment creating a ReplicaSet? Is the Scheduler scheduling?
4.  **Pod Level:** Status (`Pending`, `CrashLoopBackOff`, `ImagePullBackOff`)? Events?
5.  **Application Level:** App logs, health probes, configuration.

---

## Deep Dive: On-Prem & Control Plane Components

In managed K8s (EKS/GKE), the control plane is a black box. In on-prem/interviews, you own it.

### 1. Etcd (The Brain)
*   **Role:** Strong consistency (RAFT), single source of truth.
*   **On-Prem Risks:** Disk latency (fsync) is the #1 killer. If disk write latency > 10ms, the cluster destabilizes.
*   **Troubleshooting:**
    *   `etcdctl endpoint status --write-out=table`: Check leadership and raft index.
    *   **"Database space exceeded":** Etcd has a 2GB (default) limit. History compaction and defrag (`etcdctl defrag`) are critical maintenance.
    *   **Quorum Loss:** If you lose majority (2/3), the cluster becomes read-only. Recovery involves stopping all members and restoring from snapshot on one node (force-new-cluster).

### 2. Networking (CNI & Kube-Proxy)
*   **The Packet Flow:** Pod A -> Service B IP -> Kube-Proxy (iptables/IPVS) -> DNAT to Pod B IP.
*   **On-Prem CNI (Calico/Flannel/Cilium):**
    *   **BGP (Calico):** Peering with Top-of-Rack switches. Debug with `calicoctl node status` and `gobgp`.
    *   **IPAM:** Running out of Pod CIDRs on a node. Check `/var/lib/cni/networks`.
    *   **MTU Issues:** Overlay networks (VXLAN) add headers. If physical MTU is 1500 and VXLAN is 1550 -> Packet drops.
*   **Service Discovery (CoreDNS):**
    *   `ndots:5` issue: High DNS query volume because `musl`/`glibc` tries to resolve `name.ns.svc.cluster.local`, `name.ns.svc`, etc.
    *   Check `kubectl get endpoints` to verify the Service actually selects Pods.

### 3. The Kubelet
*   **Role:** The agent on every node. Talks to Container Runtime (CRI).
*   **Troubleshooting:**
    *   `journalctl -u kubelet`: The first place to look for "Node NotReady".
    *   **PLEG (Pod Lifecycle Event Generator) is not healthy:** The container runtime is slow/hanging. Kubelet can't update pod status.
    *   **Certificate Expiry:** Kubelet client certs expire. Bootstrap process failed?

### 4. Storage (CSI)
*   **PV vs PVC:** PVC is a request, PV is the resource. `StorageClass` is the provisioner.
*   **Common Failures:**
    *   **"VolumeAttachment" issues:** Node A dies, Node B tries to attach. Cloud provider/SAN thinks volume is still attached to Node A. Manual detach often needed.
    *   **Stuck Terminating:** Often due to `kubernetes.io/pvc-protection` finalizer. Pod is still using the volume.

---

## Common Failure Scenarios (The "What do you do?" questions)

### Scenario A: `CrashLoopBackOff`
*   **Meaning:** Process starts, crashes, restarts.
*   **Debug Steps:**
    1.  `kubectl logs <pod>` (and `kubectl logs <pod> --previous`).
    2.  Check Exit Code:
        *   `0`: App exited normally (but shouldn't have).
        *   `1`: App error.
        *   `137` (OOMKilled): `128 + 9 (SIGKILL)`. Check `kubectl describe pod` -> "Last State: OOMKilled". **Fix:** Increase memory limit or fix memory leak.
        *   `143`: Graceful termination (SIGTERM).
    3.  Check Liveness Probes: Is the app "running" but failing the health check?

### Scenario B: `Pending`
*   **Meaning:** Scheduler cannot find a node.
*   **Debug Steps:**
    1.  `kubectl describe pod`: Look at "Events".
    2.  **Insufficient CPU/Mem:** Cluster is full.
    3.  **Taints/Tolerations:** Node has `NoSchedule` taint, Pod lacks toleration.
    4.  **Affinity:** `nodeAffinity` rules cannot be satisfied.
    5.  **PVC Pending:** The storage backend can't provision the disk.

### Scenario C: Service Unreachable
*   **Meaning:** `curl service-ip` fails.
*   **Debug Steps:**
    1.  **Does the Service exist?** `kubectl get svc`.
    2.  **Are there Endpoints?** `kubectl get ep <svc-name>`. If empty, the `selector` matches no pods (check labels!).
    3.  **Is CoreDNS working?** `nslookup <svc-name>` from inside a pod.
    4.  **NetworkPolicy:** Is traffic blocked? (Default deny?).

### Scenario D: `Terminating` (Stuck)
*   **Meaning:** Pod won't die.
*   **Root Cause:** Finalizers. A resource (like a PVC or LoadBalancer) needs cleanup before the object is deleted.
*   **The "Nuclear" Option (Use with caution):** `kubectl delete pod <name> --grace-period=0 --force` (Bypasses graceful shutdown, but doesn't fix the underlying controller issue).

---

## "Hard" Interview Questions

1.  **"How do you debug a slow cluster?"**
    *   Check API Server latency (metrics).
    *   Check Etcd disk fsync latency.
    *   Check for "thundering herd" (too many controllers listing everything).
    *   Check DNS latency.

2.  **"A Node is 'NotReady'. Walk me through the debug."**
    *   `kubectl describe node`: Check conditions (DiskPressure, PIDPressure).
    *   SSH into node.
    *   Check `systemctl status kubelet`.
    *   Check `docker ps` / `crictl ps` (Runtime hanging?).
    *   Check disk space (`df -h`).

3.  **"We updated the deployment image, but the new pods aren't starting. Why?"**
    *   `maxUnavailable` setting in `RollingUpdate` strategy?
    *   `minReadySeconds` delaying rollout?
    *   Resource Quota in the namespace exceeded?

---

**Next:** [Go back to Roadmap](00-roadmap.md)
