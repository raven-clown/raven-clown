# Delta Electronics Thailand

**Application Engineer**, IT manufacturing systems team · April 2026 to present

## What the role covers

DevOps, full-stack web development, and production line support for the factory's MES (Manufacturing Execution System) tooling, inside a fully air-gapped corporate network. Also acts as project manager on almost every initiative below: planning the work, coordinating with other departments to keep it moving, presenting it to the department, and supporting the people who use it afterward.

## What I built

**Production line support.** Troubleshoot MES, application, and data issues reported from the floor, and work with the people on the line until each one is resolved.

**CI/CD platform.** Built and maintain the GitLab CI/CD pipelines for `mfg-portal-frontend` (React/Vite) and `mfg-portal-api` (ASP.NET Core), deploying through a Rancher-managed Kubernetes cluster with images pushed to Harbor. Was the first person on the Thai team to put this Rancher/Kubernetes/GitLab CI-CD architecture together and present it to the whole IT department. Other application teams and the server team are now adopting it, and I'm the point of contact for that rollout.

**MES monitoring dashboard.** Full-stack dashboard (React, .NET 8, Docker Swarm) showing real-time service health and usage metrics.

**Yield Tracker Report.** Sole developer on a full-stack production yield reporting and drill-down tool for the QA team (React/TypeScript/Tailwind, ASP.NET Core/.NET 8, Oracle via ODP.NET, SQL Server holding query templates as config-as-data). Aggregates station-level pass/fail data across factories and models, with drill-down from summary down to individual repair records and Excel export via ExcelJS. Fixed an N+1 query in MO-to-model resolution, parallelized per-schema data fetching, added row virtualization for 10,000+ row tables without pagination, and traced a session-hydration race condition in the app's shared auth layer that had been silently breaking deep links app-wide. Deployed in Thailand with rollout planned for Taiwan and other sites; took about five days to build.

**CTP diagnostic tool.** ASP.NET Core API + React/TypeScript frontend that automates cross-schema comparisons in Oracle (`PROD_SCHEMA_A` / `PROD_SCHEMA_B`), flagging MATCH/MISMATCH cases that used to be checked manually.

**MinIO object storage.** Set up and configured a MinIO deployment from scratch and advise other teams on how to deploy against it. Also deployed a MinIO instance as a Docker container to back the request review system below.

**Request review system.** Designed and built the backend for a request system that collects the PowerPoint files, documents, and images people submit, and plugs into review flows that already exist, such as MES requests and VM server requests that have to pass review before they're approved. A Java Spring Boot backend exposes the REST API the web frontend calls to submit, upload, and review requests; submitted files go to MinIO, and request data, work history, and review records live in Microsoft SQL Server.

**AIFinance.** Led a project for the finance side: a system that pulls finance documents every month, matches related records across sources, and uses AI to produce summaries and analysis on vendors and customers, covering currency and salary data. Started as a test application that ran existing data through the backend to check the AI's analysis before anything was built for users, then moved to a web app that keeps historical data for trend analysis and refines the AI's output against real documents. Python and FastAPI backend, PostgreSQL, a React frontend, and email notifications for results and status. Coordinated with the stakeholders and designed the architecture as project manager.

**Central NiFi platform.** Set up and run a dedicated NiFi cluster, separate from the team's shared production clusters, as a central platform that IT and other teams build their own flows on. Used as middleware for logging, alerting, data migration, and routing data across multiple systems, including running scripts for MinIO job work. Presented this to the department as well.

**AIDeltron log analytics platform.** Built a platform that collects logs from the factory's applications and machines into OpenSearch and turns them into something people can investigate. Python ingestion connects to sources across many IP addresses and pulls `.log` files automatically from sources that each use a different format (API logs in several shapes, routing logs, SMT machine logs, aggregator logs from production equipment), converts them to NDJSON, and bulk-loads them into a two-node OpenSearch cluster, either on a configurable schedule or through manual upload. A FastAPI backend sits between the UI and OpenSearch. Storage is split by weight: SQLite for lightweight data, and PostgreSQL for anything large or critical (which files have already been fetched, aggregated summaries, history, an audit log of every action, and AI token usage monitoring), with old logs archived automatically. The web UI has a machine log view, search across all logs, crosstab views, and per-source fetch settings, and it calculates daily issue counts and compares them against targets and historical values to track trends. Searching a log shows its full detail along with the likely cause, plus charts, AI summaries, and an AI assistant that answers questions and helps people work back through log history. Every source can be enabled, disabled, fetched, checked, and configured from the UI without touching the server. An MCP server on top of the same data lets an agent ask questions, compare patterns across sources, and trace the root cause of an application or machine failure.

**Central Kafka platform.** Built a shared Kafka platform that IT and other application teams can plug their own apps into, instead of each team standing up its own broker. A 3-broker cluster in KRaft mode (no ZooKeeper, replication factor 3, every node running broker, controller and JMX exporter), plus a separate tooling VM with AKHQ for topic and consumer lag inspection, Prometheus scraping each broker's JMX exporter, Grafana dashboards, and Redis for consumer services, all behind an nginx reverse proxy. Also specified the broker sizing (4 vCPU, 8 GB RAM, 50 GB OS + 300 GB data disk, RHEL 8.10, non-root) and the full port plan for the servers.

**Data pipeline design.** Separately from the Kafka and NiFi platforms themselves, designed the flows that pull data out and send it on: NiFi reads from the Oracle MES databases over CDC/JDBC and produces raw events into Kafka, Spark Structured Streaming on the existing data processing cluster consumes the raw topics and writes cleaned gold topics, and Airflow workflows schedule the Spark jobs every 15 to 30 minutes.

**Data investigation.** Traced a `shipment_code` vs internal `unit_serial` mismatch across the `PROD_SCHEMA_B` schema to support a live MES rework case.

**Internal knowledge search tool.** Built a retrieval-based document search tool for the team. It syncs and keeps up to date documents from SharePoint and Confluence, along with past ticket responses and resolutions from ServiceDesk Plus (SDP). Documents are chunked ahead of time and embedded with Qwen embedding models, with the Qwen API used as the underlying LLM to answer queries grounded in the retrieved text. Every conversation is logged to build datasets for continuously fine-tuning the model. Extended beyond plain text retrieval with reranking, vision-language models for reading scanned/image-based documents, and speech-to-text for audio sources, plus using logs and API/error data indexed in OpenSearch as fine-tuning input for smaller models trained on operational data directly.

**Observability stack.** Set up Prometheus and Grafana for metrics, and OpenSearch for centralized log storage and search across the team's services.

## Full technical write-up

See [`brain/devops-infrastructure`](../../brain/devops-infrastructure), [`brain/backend-development`](../../brain/backend-development), and [`brain/ai-integration`](../../brain/ai-integration) for the underlying technical detail.
