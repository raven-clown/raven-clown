# DevOps & Infrastructure

How the pieces of a CI/CD and deployment stack work, and how they fit together.

## Architecture — A Typical Pipeline

```mermaid
flowchart LR
    Dev["Commit"] --> CI["CI/CD pipeline"]
    CI --> Build["Build & test"]
    Build --> Registry["Push image\nto container registry"]
    Registry --> Deploy["Authenticated deploy\n(API token)"]
    Deploy --> K8s["Kubernetes cluster"]
    K8s --> FE["Frontend service"]
    K8s --> BE["Backend service"]
    K8s --> Storage[("Object storage")]
```

## GitLab CI/CD

- A pipeline is defined in a `.gitlab-ci.yml` file and runs as a sequence of stages — typically build, test, then deploy — where each stage only starts once the previous one passes.
- A runner (the "executor") is the machine that actually executes each job. A shell executor runs jobs directly on the host's shell instead of spinning up a fresh container per job — faster to configure, but the environment has to be kept clean manually since it isn't reset between runs.
- Deployment jobs authenticate against a target cluster using a scoped API token rather than a personal login, so access can be revoked or rotated without touching anyone's individual credentials.
- Artifacts (build output, test reports) can be passed between stages so a later job doesn't have to rebuild what an earlier one already produced.

## Kubernetes & Rancher

- Kubernetes schedules containers ("pods") across a cluster of machines, and keeps the actual state matching whatever you declared (a Deployment says "I want 3 replicas of this pod running" — Kubernetes keeps making that true).
- Namespaces partition a cluster into logical environments (dev, QA, prod, or per-team) that share the underlying hardware but stay isolated from each other for access control and resource limits.
- Rancher sits on top of raw Kubernetes as a management layer — one UI/API for multiple clusters, centralized RBAC, and easier day-to-day operations than talking to each cluster's API directly.
- A push-based deploy (`kubectl apply` from CI, authenticated with an API token) is simpler to set up than a pull-based GitOps model, at the cost of the cluster having to trust an external system to push changes to it.
- Ingress controllers handle routing external traffic into the right service inside the cluster based on hostname/path rules, instead of exposing every service on its own IP or port.
- A Horizontal Pod Autoscaler (HPA) scales replica count based on load, but it's still bound by whatever ResourceQuota is set on the namespace — if the quota is maxed out, the HPA can be "correctly" trying to scale up and still get blocked.

## Docker & Docker Swarm

- An image is built in layers; BuildKit (Docker's newer build engine) caches those layers intelligently so unchanged layers don't get rebuilt on every run.
- Tagging images with the commit SHA instead of a static tag like `latest` means every deployed image is traceable back to the exact commit that produced it — useful when something breaks and you need to know exactly what's running.
- Docker Swarm is a simpler orchestrator than Kubernetes — fewer moving parts, easier to reason about for smaller deployments, at the cost of a smaller feature set (no equivalent to Kubernetes' Ingress/HPA ecosystem out of the box).

## Harbor (Container Registry)

- A private registry stores built images so they don't have to be pulled from a public registry — important in restricted network environments, and it gives you control over retention and vulnerability scanning.
- Base images (like `nginx:alpine`) need to be deliberately preserved/pinned in a private registry, or a routine cleanup policy can delete an image that's still in active use by a running deployment.

## Apache NiFi

- A dataflow automation tool built around a visual pipeline: each step (a "processor") reads, transforms, routes, or writes data, and processors are wired together into a flow instead of hand-writing point-to-point integration scripts.
- Works well as a general middleware layer — logging, alerting, data migration, and routing data between systems can all live as flows in the same tool, so a new integration is usually a new flow rather than a new one-off script.
- Flows are visual and inspectable at runtime — a stuck or failing step shows up directly in the flow itself, which makes tracing where a pipeline broke faster than digging through separate log files for each system it touches.

## MinIO (Object Storage)

- An S3-compatible object store — same API shape as AWS S3, but self-hosted. Anything already written against the S3 SDK works against it with just an endpoint change.
- Organized into buckets with their own access policies, which makes it a natural drop point for pipeline artifacts, backups, or files that a batch job needs to pick up and process.

## AWS

- Running a Kubernetes cluster on cloud infrastructure means the control plane and worker nodes are cloud-managed resources instead of physical machines, which makes scaling a capacity decision instead of a hardware-ordering one.
- A typical setup: a load balancer or CDN in front, SSL terminated at the edge, then traffic routed back into the cluster.

## Secrets Management

- Symmetric encryption (AES-256-CBC) is a reasonable way to store secrets in a config file that has to live in a repo — the encrypted value is safe to commit, and only something holding the key can decrypt it back.
- If a real secret ever lands in git history unencrypted, rotating the secret alone isn't enough — the history itself has to be rewritten (with a tool like `git filter-repo`) or the old value stays recoverable by anyone who already has a clone.
