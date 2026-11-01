# System Architecture

Diagrams of the systems described across this repository. The front page keeps the summary; this page keeps the drawings.

### [ARK](https://github.com/raven-clown/ark): From a Topic to Every System That Needs the Message

One engine per node, coordinated through Kafka, with a console and agents on the side.

```mermaid
flowchart LR
    Src[("Kafka topic")] --> Engine["ARK engine\n(Go, one binary)"]
    Engine --> Check{"data rules\nand conditions"}
    Check -->|"passes"| App["Your app\nover HTTP"]
    Check -->|"breaks a rule"| Rej[("reject topic")]
    App -->|"2xx"| Res[("result topic")]
    App -->|"2xx"| Out["OpenSearch / NiFi /\nSlack / webhooks"]
    App -.->|"retries used up"| DLQ[("dead-letter topic")]
    Console["ARK Console\n(React + React Flow)"] <--> Engine
    Agents["MCP agents\nand Ask ARK"] <--> Engine
    Engine <-.->|"heartbeats, leader,\nshared config"| Nodes["Other ARK nodes"]
```

### [tack](https://github.com/raven-clown/tack): One Struct, Every Output

Everything is derived at registration from the handler's types; the request path runs precomputed plans only, and every generated artifact reads the same OpenAPI document.

```mermaid
flowchart LR
    Structs["Handler In / Out structs\npath, query, header, cookie,\nform and Body tags"] --> Reg["Registration, once\nbind plan, schema registry,\nOpenAPI 3.1 document"]
    Reg --> Req["Request time\nplan, validator, handler"]
    Req --> Resp["JSON, problem+json,\nserver-sent events"]
    Reg --> Doc[("OpenAPI document")]
    Doc --> Docs["/docs"]
    Doc --> TS["api.ts\nResult unions by error code"]
    Doc --> Zod["schemas.ts\nZod with the same rules"]
    Doc --> Mock["tack mock"]
    Dev["tack dev\nrebuild and regenerate on save"] -.-> Doc
    Check["tack check / tsgen.Check\nCI fails on a stale client"] -.-> TS
```

### Delta MES: CI/CD and Deployment

How code gets from a commit to a running service inside the air-gapped cluster.

```mermaid
flowchart LR
    Dev["Commit"] --> CI["GitLab CI/CD\n(shell executor)"]
    CI --> Build["Build & test"]
    Build --> Harbor["Push image\nto Harbor registry"]
    Harbor --> Deploy["kubectl deploy\n(Rancher API token)"]
    Deploy --> K8s["Kubernetes cluster\nmfg-cluster-dev"]
    K8s --> FE["mfg-portal-frontend\n(React / Vite)"]
    K8s --> BE["mfg-portal-api\n(ASP.NET Core)"]
    K8s --> MinIO[("MinIO\nobject storage")]
    BE --> Oracle[("Oracle DB\nPROD_SCHEMA_A / PROD_SCHEMA_B")]
```

### Delta MES: CTP Diagnostic Tool, Yield Tracker, and Dashboard

Where the data actually comes from and how it reaches the screen.

```mermaid
flowchart LR
    Oracle[("Oracle DB\nPROD_SCHEMA_A / PROD_SCHEMA_B")] --> API["ASP.NET Core API"]
    MSSQL[("MS SQL Server")] --> API
    PG[("PostgreSQL")] --> API
    API --> CTP["CTP comparison engine\n(CTE-based MATCH / MISMATCH)"]
    API --> Yield["Yield Tracker engine\n(station pass/fail aggregation)"]
    API --> Metrics["/api/health\n/api/metrics"]
    CTP --> UI1["CTP diagnostic UI\n(React / TypeScript)"]
    Yield --> UI3["Yield Tracker UI\n(React / TypeScript / Tailwind)"]
    Metrics --> UI2["MES dashboard\n(React + .NET 8 + Docker Swarm)"]
```

### Delta: AI OpenSearch Log Analytics

Only the new bytes of every log, parsed, searchable, and turned into Teams alerts while production is still affected.

```mermaid
flowchart LR
    Src["Gateway / routing / SMT /\nIoT aggregator log servers"] -->|"new bytes only\n(HTTP Range)"| Col["Collector\n(Go)"]
    Col --> Jobs[("PostgreSQL\njob queue")]
    Jobs --> W["Workers (Go, auto-scaled)\nrunning Python parsers"]
    W --> Imp["Importer\n(Go)"]
    Imp --> OS[("OpenSearch\ndaily indices")]
    OS --> API["FastAPI\nsearch, issues,\nAI assistant, exports"]
    OS --> Watch["Alert watcher\n(own container)"]
    Watch -->|"outbox, retries"| Teams["Teams\nwebhook or channel email"]
    Nginx["nginx\nblue/green API"] --> API
    API --> UI["MES portal page and\nAI OpenSearch page"]
    API --> Ext["External API\nfor other AI tools"]
```

### Delta: AI Center, AI Deltron and AI Monitor

AI Center is the web portal for every AI project the team has built. AI Deltron answers from the company's own knowledge, and AI Monitor is the admin console behind both. The backend was split into services one step at a time.

```mermaid
flowchart LR
    User["Browser\nSSO login"] --> Nginx["nginx"]
    Nginx --> Portal["AI Center\n(React 19)"]
    Nginx --> Edge["Edge\n(Go)"]
    Edge --> API["API role"]
    Edge --> Chat["AI Deltron chat\nretrieval answers"]
    Edge --> Media["Media role\nstaff photos as WebP"]
    Worker["Worker role\nConfluence / SharePoint sync"] --> PG
    API --> PG[("Postgres + pgvector")]
    Chat --> PG
    Chat --> AIGW["AI gateway proxy (Go)\nquota share, 429 queue"]
    AIGW --> Corp["Company AI gateway"]
    API --> Store[("MinIO / NATS")]
    Ops["Ops (Go)\ncontainer scaling"] -.-> API
    Analytics["Analytics (Go)\nAI Monitor numbers"] --> PG
```

### Delta: AIFinance

Company matching for Finance: the same company found across SAP exports, put under its group, with every number from code.

```mermaid
flowchart LR
    Files["Monthly SAP exports\nmail or upload"] --> API["Web backend\n(FastAPI, roles)"]
    API --> Queue[("NATS JetStream\nrun queue")]
    Queue --> Worker["Workers\nread, normalize, exclude,\nmatch, group, check, summarize"]
    Worker --> GW["AI gateway client\npaced batch questions"]
    Know[("Knowledge file\noutranks the AI")] --> Worker
    Worker --> PG[("PostgreSQL\nstored decisions, runs")]
    PG --> UI["Dashboard\nreview and release"]
    UI --> Export["xlsx / pdf / docx / pptx\nfour languages"]
```

### Delta: Requirement Request Tracking

Requests that move through MES, AP, IT Server and OA review, stored as append-only snapshots.

```mermaid
flowchart LR
    Web["React frontend"] -->|"JWT"| API["FastAPI backend\nworkflow rules"]
    API -->|"SELECT / INSERT only\nsnapshot per change,\napp lock + revision check"| SQL[("SQL Server")]
    API -->|"files, metadata,\ncommit markers"| MinIO[("MinIO")]
    API --> Flow["MES / AP / IT Server / OA\nreview stages"]
```

### Delta: Central Kafka and NiFi Platforms

Shared platforms other IT teams plug their apps into, and the data flows built on top of them.

```mermaid
flowchart LR
    Oracle[("Oracle MES DBs")] -->|"CDC / JDBC"| NiFi["Central NiFi\nplatform"]
    NiFi -->|"raw topics"| Kafka[("Central Kafka platform\nKRaft, 3 brokers, RF 3")]
    Apps["Other IT apps"] <--> Kafka
    Kafka --> Spark["Spark Structured\nStreaming"]
    Spark -->|"gold topics"| Kafka
    Airflow["Airflow workflows\nevery 15-30 min"] -.-> Spark
    Kafka -->|"JMX"| Prom["Prometheus / Grafana"]
    Nginx["nginx"] --> AKHQ["AKHQ"] --> Kafka
```

### FiveM Server Platform

Server, systems, and the storefront running alongside it.

```mermaid
flowchart TB
    Client["Game client"] <--> Server["Windows Server\n(i9, 64GB RAM, 10Gbps)"]
    Nginx["Nginx CDN cache"] --> Client
    Server --> Core["raven_core\neconomy / identity, batched writes"]
    Server --> Medical["raven_advanceambulance\nknockdown / crawl state machine"]
    Server --> Pedscale["raven_pedscale"]
    Server --> Opt["myopt\ncustom 7-module optimizer"]
    Core --> DB[("MySQL\nvia oxmysql")]
    Server --> UI["HUD / inventory UI\n(Svelte + React)"]
    Server --> Shop["shop.highclass-roleplay.com"]
    Shop --> Cloud["Next.js + Go\non AWS / Kubernetes"]
```

### [IdpForge](https://github.com/raven-clown/idpforge): Cluster-Safe Background Jobs and Realtime Fan-Out

How one binary stays correct whether it's a single instance or several behind a load balancer.

```mermaid
flowchart LR
    Browser["Browser\n(embedded Next.js SPA)"] --> API["Go API\n(single binary, N instances)"]
    API --> RBAC["RBAC resolver\n(cached, invalidated on write)"]
    API --> OIDC["OIDC provider\nauth code+PKCE, client_credentials"]
    API --> Hub["Realtime hub\n(WebSocket)"]
    API --> Lease["Leader lease\n(delete-expired, insert, UPDATE by PK)"]
    RBAC --> DB[("Postgres / MySQL /\nMSSQL / SQLite")]
    Lease --> DB
    Hub -. "Redis pub/sub" .-> Other["Every other instance's\nconnected browsers"]
```

### Snippet Manager: Two Front-Ends, One Core

How the VS Code extension and the Electron app share state without a server.

```mermaid
flowchart LR
    Core["@snippet/core\nstorage, CRUD, validation, migration"]
    Core <--> DataFile[("data.json\nOS app-data directory")]
    VSC["vscode-extension"] --> Core
    Electron["electron-app"] --> Core
    DataFile -. "chokidar watch" .-> VSC
    DataFile -. "chokidar watch" .-> Electron
```

### [RoomedIn](https://www.roomedin.online/): Defense in Depth on a Status Change

Five independent layers between a tap on a phone and a row actually changing.

```mermaid
flowchart LR
    UI["Staff taps\nnext status"] --> Domain["Domain\nRoom.changeStatus()"]
    Domain --> App["Application\nactor/hotel check"]
    App --> MW["Next.js middleware\nrole gating"]
    MW --> RLS[("Postgres RLS\ncolumn-level grants")]
    RLS --> Trigger["DB trigger\nwrites audit log"]
    RLS --> Realtime["Supabase Realtime"]
    Realtime --> Board["Every other\nstaff device"]
```
