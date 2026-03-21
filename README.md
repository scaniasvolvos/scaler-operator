# Scaler Operator 🚀

A simple Kubernetes Operator built using Operator SDK that performs **time-based scaling of Deployments**.

---

## 📌 Overview

Scaler Operator allows you to automatically scale one or more Kubernetes Deployments to a desired number of replicas **within a specified time window (UTC)**.

This is useful for:

* Cost optimization (scale down during off-hours)
* Handling predictable traffic spikes
* Scheduled scaling without relying on HPA

---

## ⚙️ How It Works

1. Define a custom resource (`Scaler`)
2. Specify:

   * Start and end time (UTC hours)
   * Desired replicas
   * Target deployments
3. The controller:

   * Runs every 30 seconds
   * Checks current time
   * Scales deployments if within the defined window

---

## 📦 Custom Resource Example

```yaml
apiVersion: api.dpranav.online/v1alpha1
kind: Scaler
metadata:
  labels:
    app.kubernetes.io/name: scaler-operator
    app.kubernetes.io/managed-by: kustomize
  name: scaler-sample
spec:
  start: 5        # 5 AM UTC
  end: 7          # 7 AM UTC
  replicas: 5
  deployments:
    - name: nginx
      namespace: default
```

---

## 🧱 Project Structure

```
.
├── api/v1alpha1/               # CRD definitions (Spec & Status)
├── internal/controller/        # Reconciliation logic
├── config/                     # Kubernetes manifests (CRDs, RBAC, samples)
├── cmd/main.go                 # Entry point
├── Dockerfile
├── Makefile
└── PROJECT                     # Kubebuilder project metadata
```

---

## 🚀 Getting Started

### Prerequisites

* Go (>= 1.24)
* Docker
* kubectl
* Operator SDK
* Kubernetes cluster (Minikube / Kind / EKS etc.)

---

### 🛠️ Build and Deploy

```bash
# Generate manifests
make manifests

# Build docker image
make docker-build IMG=<your-docker-image>

# Push image
make docker-push IMG=<your-docker-image>

# Deploy CRDs and controller
make deploy IMG=<your-docker-image>
```

> **Note:** If you encounter RBAC errors, you may need to grant yourself cluster-admin privileges or be logged in as admin.

---

### 📥 Apply a Scaler Resource

```bash
kubectl apply -f config/samples/
```

---

### 🗑️ Uninstall

```bash
# Delete the Scaler CRs
kubectl delete -k config/samples/

# Delete the CRDs
make uninstall

# Remove the controller
make undeploy
```

---

## 🔄 Reconciliation Behavior

* Runs every **30 seconds**
* Uses **UTC time**
* Scales deployments only within the defined window
* Updates status as:

  * `Success`
  * `Failed`

---

## ⚠️ Limitations

* Timezone not configurable (UTC only)
* No scale-down logic outside time window
* Basic status reporting
* No metrics or observability

---

## 💡 Future Improvements

* Add timezone support
* Cron-based scheduling
* Integration with metrics (CPU-based scaling)
* Better error handling and retries
* Observability (Prometheus metrics)

---

## 📜 License

Copyright 2026 Pranav Deshpande.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

---

## 🙌 Acknowledgements

Built using:

* Operator SDK
* Kubebuilder
* Kubernetes controller-runtime

---
