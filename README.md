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
apiVersion: api.adityajoshi.online/v1alpha1
kind: Scaler
metadata:
  name: sample-scaler
spec:
  start: 9
  end: 18
  replicas: 5
  deployments:
    - name: nginx
      namespace: default
```

---

## 🧱 Project Structure

```
.
├── api/v1alpha1        # CRD definitions (Spec & Status)
├── controllers         # Reconciliation logic
├── config              # Kubernetes manifests
├── main.go             # Entry point
```

---

## 🚀 Getting Started

### Prerequisites

* Go (>= 1.19)
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

---

### 📥 Apply a Scaler Resource

```bash
kubectl apply -f config/samples/
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

Licensed under the Apache License 2.0.

---

## 🙌 Acknowledgements

Built using:

* Operator SDK
* Kubebuilder
* Kubernetes controller-runtime

---
