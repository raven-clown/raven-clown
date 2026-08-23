<div align="center">

# Ekdanai Kummee (RAVEN)

**Application Engineer**, Delta Electronics Thailand

[GitHub: raven-clown](https://github.com/raven-clown)

<br>

[![My Skills](https://skillicons.dev/icons?i=discord,jquery,cs,ts,js,go,lua,py,java,php,html,css,react,nextjs,nuxtjs,vue,svelte,dotnet,express,nestjs,electron,nodejs,vite,redis,bun,figma,git,gitlab,docker,kubernetes,aws,vercel,postgres,mysql,mongodb,supabase,arduino,linux,vscode,visualstudio&perline=10)](https://skillicons.dev)

<br>

![Oracle](https://img.shields.io/badge/Oracle-F80000?style=flat-square&logo=oracle&logoColor=white)
![Microsoft SQL Server](https://img.shields.io/badge/SQL_Server-CC2927?style=flat-square&logo=microsoftsqlserver&logoColor=white)
![MinIO](https://img.shields.io/badge/MinIO-C72E49?style=flat-square&logo=minio&logoColor=white)
![Harbor](https://img.shields.io/badge/Harbor-60B932?style=flat-square)
![Rancher](https://img.shields.io/badge/Rancher-0075A8?style=flat-square)
![LINE](https://img.shields.io/badge/LINE_API-06C755?style=flat-square&logo=line&logoColor=white)
![Cisco](https://img.shields.io/badge/Cisco-1BA0D7?style=flat-square&logo=cisco&logoColor=white)
![Power BI](https://img.shields.io/badge/Power_BI-F2C811?style=flat-square&logo=powerbi&logoColor=black)
![Machine Learning](https://img.shields.io/badge/Machine_Learning-FF6F00?style=flat-square)
![Red Hat Enterprise Linux](https://img.shields.io/badge/RHEL-EE0000?style=flat-square&logo=redhat&logoColor=white)

<br>

![Skill Proficiency](./assets/skills-chart.svg)

<br>

![Language Proficiency](./assets/languages-chart.svg)

</div>

---

## About

I'm a junior developer at Delta Electronics Thailand, on the IT manufacturing systems team, where I handle both DevOps and full-stack web development for the factory's MES (Manufacturing Execution System) tooling. I also run my own projects end to end: planning them, building them, presenting them to the team, and supporting the people who use them afterward.

Outside of work, I've spent about five years building and operating multiplayer game-server platforms (FiveM and RedM), which is where most of my backend, database, and systems-architecture experience actually comes from before I ever touched a corporate codebase. I taught myself Lua well enough to write almost anything in it from memory, and that same instinct for "just build the thing and fix what breaks" carried over into everything else here: DevOps, backend APIs, database design, and UI.

This repository is a map of that work. This page is the summary.

- [`about/`](./about): who I am outside of the job title, how I work, and personal background.
- [`experience/`](./experience): the portfolio, concrete projects and what came out of them.
- [`brain/`](./brain): the knowledge base, how each skill area actually works, in detail.

**Education:** B.Sc. Digital Business Technology, Southeast Bangkok University (2025, GPA 3.63, First Class Honors / Gold Medal) · High Vocational Certificate in Information Technology, Attawit Commercial Technology College (2023, GPA 3.10)

**Languages:** Thai (native), English (working proficiency)

---

## What I Actually Do

Each area below is split two ways. **Knowledge** explains how the technology itself works. **Portfolio** shows what I actually built with it. They're kept separate on purpose: one is reference material, the other is proof of work.

### DevOps & Infrastructure

CI/CD pipelines, Kubernetes/Rancher, Docker, container registries, object storage, and cloud hosting. The plumbing that gets code from a commit to a running service.

Knowledge: [`brain/devops-infrastructure`](./brain/devops-infrastructure) · Portfolio: [`experience/delta-electronics`](./experience/delta-electronics)

### Backend Development

ASP.NET Core, NestJS with Prisma, and Go. How each one structures a request, handles data access, and where they're each the right tool.

Knowledge: [`brain/backend-development`](./brain/backend-development) · Portfolio: [`experience/delta-electronics`](./experience/delta-electronics) · [`experience/side-projects`](./experience/side-projects)

### Frontend Development

React, Next.js, Nuxt.js, Vue, and Svelte, plus UX/UI design for real-time interfaces where the UI can't get in the way of what's happening underneath it.

Knowledge: [`brain/frontend-development`](./brain/frontend-development) · Portfolio: [`experience/delta-electronics`](./experience/delta-electronics) · [`experience/fivem-redm-platform`](./experience/fivem-redm-platform) · [`experience/side-projects`](./experience/side-projects)

### Database Design

Oracle, SQL Server, PostgreSQL, MySQL, MongoDB, and Supabase. Relational modeling, indexing, and schema design for both enterprise data and high-write real-time systems.

Knowledge: [`brain/database-design`](./brain/database-design) · Portfolio: [`experience/delta-electronics`](./experience/delta-electronics) · [`experience/fivem-redm-platform`](./experience/fivem-redm-platform) · [`experience/side-projects`](./experience/side-projects)

### FiveM & RedM Development

Five years of building on the FiveM and RedM multiplayer platforms. Server architecture, the Lua resource model, and the ESX/QBCore frameworks that sit underneath most roleplay servers.

Knowledge: [`brain/fivem-redm-development`](./brain/fivem-redm-development) · Portfolio: [`experience/fivem-redm-platform`](./experience/fivem-redm-platform)

### AI / LLM Integration

Retrieval-based systems that ground an LLM's answers in real documents instead of relying on what the model already knows. Chunking, retrieval, and prompt construction.

Knowledge: [`brain/ai-integration`](./brain/ai-integration) · Portfolio: [`experience/delta-electronics`](./experience/delta-electronics)

### Discord Bots & Automation

Bot architecture, job queues, and the runtimes/libraries that make background automation reliable instead of flaky.

Knowledge: [`brain/discord-bots-automation`](./brain/discord-bots-automation) · Portfolio: [`experience/fivem-redm-platform`](./experience/fivem-redm-platform) · [`experience/side-projects`](./experience/side-projects)

### Embedded & Hardware

Arduino programming, sensor input, and circuit prototyping. Where software meets a physical breadboard.

Knowledge: [`brain/embedded-hardware`](./brain/embedded-hardware) · Portfolio: [`experience/academic-projects`](./experience/academic-projects)

### System Design Patterns

How I approach scoping an architecture to an actual budget, structuring a multi-service system, and deciding what to build versus what to borrow.

Knowledge: [`brain/side-projects-architecture`](./brain/side-projects-architecture) · Portfolio: [`experience/side-projects`](./experience/side-projects)

### Tools & Troubleshooting

Environment debugging, git workflows, and the unglamorous work that has to happen before anything above ships.

Knowledge: [`brain/tools-and-troubleshooting`](./brain/tools-and-troubleshooting)

---

## System Architecture

A few of the systems above, laid out.

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
    API --> CTP["CTP comparison engine\n(CTE-based MATCH / MISMATCH)"]
    API --> Yield["Yield Tracker engine\n(station pass/fail aggregation)"]
    API --> Metrics["/api/health\n/api/metrics"]
    CTP --> UI1["CTP diagnostic UI\n(React / TypeScript)"]
    Yield --> UI3["Yield Tracker UI\n(React / TypeScript / Tailwind)"]
    Metrics --> UI2["MES dashboard\n(React + .NET 8 + Docker Swarm)"]
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

### RoomedIn: Defense in Depth on a Status Change

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

---

## Core Stack

**Languages:** C#, TypeScript, JavaScript, Python, Java, PHP, Go, Lua, SQL
**Frameworks:** React, Next.js, Nuxt.js, Vue.js, Svelte, Electron, ASP.NET Core (.NET 8), Express.js, NestJS
**Infrastructure:** GitLab CI/CD, Docker, Docker Swarm, Kubernetes (Rancher), Harbor Registry, AWS, Vercel, MinIO
**Databases:** Oracle, Microsoft SQL Server, PostgreSQL, MySQL, MongoDB, Supabase
**Data:** Power BI, machine learning fundamentals
**Other:** LINE Bot/OA API, Figma, Arduino and embedded systems, Cisco networking, Adobe Photoshop/Illustrator, Linux (RHEL), Visual Studio, VS Code

---

## Contact

- Email: ekdanai.kk@gmail.com
- GitHub: [raven-clown](https://github.com/raven-clown)
