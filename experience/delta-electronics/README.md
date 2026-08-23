# Delta Electronics Thailand

**Web App Engineer**, IT manufacturing systems team · April 2026 to present

## What the role covers

Both DevOps and full-stack web development for the factory's MES (Manufacturing Execution System) tooling, inside a fully air-gapped corporate network. Also acts as project manager on the initiatives below, planning, presenting to the department, and handling support afterward.

## What I built

**CI/CD platform.** Built and maintain the GitLab CI/CD pipelines for `mfg-portal-frontend` (React/Vite) and `mfg-portal-api` (ASP.NET Core), deploying through a Rancher-managed Kubernetes cluster with images pushed to Harbor. Was the first person on the Thai team to put this Rancher/Kubernetes/GitLab CI-CD architecture together and present it to the whole IT department. Other application teams and the server team are now adopting it, and I'm the point of contact for that rollout.

**MES monitoring dashboard.** Full-stack dashboard (React, .NET 8, Docker Swarm) showing real-time service health and usage metrics.

**Yield Tracker Report.** Sole developer on a full-stack production yield reporting and drill-down tool for the QA team (React/TypeScript/Tailwind, ASP.NET Core/.NET 8, Oracle via ODP.NET, SQL Server holding query templates as config-as-data). Aggregates station-level pass/fail data across factories and models, with drill-down from summary down to individual repair records and Excel export via ExcelJS. Fixed an N+1 query in MO-to-model resolution, parallelized per-schema data fetching, added row virtualization for 10,000+ row tables without pagination, and traced a session-hydration race condition in the app's shared auth layer that had been silently breaking deep links app-wide. Deployed in Thailand with rollout planned for Taiwan and other sites; took about five days to build.

**CTP diagnostic tool.** ASP.NET Core API + React/TypeScript frontend that automates cross-schema comparisons in Oracle (`PROD_SCHEMA_A` / `PROD_SCHEMA_B`), flagging MATCH/MISMATCH cases that used to be checked manually.

**MinIO object storage.** Set up and configured a MinIO deployment from scratch and advise other teams on how to deploy against it.

**NiFi middleware.** Set up and run a dedicated NiFi cluster, separate from the team's shared production clusters, used as middleware for logging, alerting, data migration, and routing data across multiple systems, including running scripts for MinIO job work. Presented this to the department as well.

**Data investigation.** Traced a `shipment_code` vs internal `unit_serial` mismatch across the `PROD_SCHEMA_B` schema to support a live MES rework case.

**Internal knowledge search tool.** Built a retrieval-based document search tool for the team. Documents chunked ahead of time, with the Qwen API used as the underlying LLM to answer queries grounded in the retrieved text.

## Full technical write-up

See [`brain/devops-infrastructure`](../../brain/devops-infrastructure), [`brain/backend-development`](../../brain/backend-development), and [`brain/ai-integration`](../../brain/ai-integration) for the underlying technical detail.
