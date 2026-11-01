<div align="center">

# Ekdanai Kummee

### Application Engineer · DevOps, data platforms and full-stack web

Building the systems behind a factory's MES at **Delta Electronics Thailand**, and open source tooling for Kafka and Go on the side.

<br>

[![GitHub](https://img.shields.io/badge/GitHub-raven--clown-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/raven-clown)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Ekdanai_Kummee-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ekdanai-kummee-793ba43ab/)
[![Email](https://img.shields.io/badge/Email-ekdanai.kk%40gmail.com-EA4335?style=for-the-badge&logo=gmail&logoColor=white)](mailto:ekdanai.kk@gmail.com)

</div>

<br>

## About

- Application Engineer on the IT manufacturing systems team at Delta Electronics Thailand: CI/CD, Kubernetes, full-stack web, and production line support for MES tooling in an air-gapped network.
- I build and run the shared data platforms other teams plug into: a central Kafka cluster, a NiFi cluster, streaming pipelines out of the MES, and a log analytics platform on OpenSearch.
- I usually own what I build end to end, from scoping and cross-department coordination to presenting it and supporting the people who use it.
- Before corporate work I spent five years founding and operating a multiplayer game-server platform (FiveM and RedM), which is where most of my backend and systems experience started.

<br>

## Tech Stack

<div align="center">

[![Skills](https://skillicons.dev/icons?i=go,cs,ts,js,py,lua,react,nextjs,svelte,tailwind,nodejs,dotnet,nestjs,fastapi&perline=14)](https://skillicons.dev)

[![Skills](https://skillicons.dev/icons?i=kafka,rabbitmq,postgres,mysql,redis,elasticsearch,docker,kubernetes,nginx,gitlab,githubactions,linux,grafana,prometheus&perline=14)](https://skillicons.dev)

</div>

| Area | Tools |
|---|---|
| **Backend** | Go, ASP.NET Core, NestJS, FastAPI, Spring Boot |
| **Frontend** | React, Next.js, Svelte, Vue, Tailwind CSS |
| **Data and messaging** | Kafka, RabbitMQ, NATS, Apache NiFi, OpenSearch, Spark Structured Streaming, Airflow |
| **Databases** | Oracle, SQL Server, PostgreSQL, MySQL, MongoDB, Redis |
| **Infrastructure** | GitLab CI/CD, GitHub Actions, Docker, Kubernetes (Rancher), Harbor, MinIO, Prometheus, Grafana, AWS |

<br>

## Featured Projects

<table>
<tr>
<td width="33%" valign="top">
<a href="https://github.com/raven-clown/ark"><img src="./project-images/covers/ark-kafkatohttp-bridge.png" alt="ARK"></a>
<h3 align="center"><a href="https://github.com/raven-clown/ark">ARK</a></h3>
<p align="center">Kafka-to-HTTP bridge in Go. Retries, dead letters, flows, cluster mode and a React console. One binary, one YAML file.</p>
</td>
<td width="33%" valign="top">
<a href="https://github.com/raven-clown/tack"><img src="./project-images/covers/tack-typesafe-http-apis-for-go.png" alt="tack"></a>
<h3 align="center"><a href="https://github.com/raven-clown/tack">tack</a></h3>
<p align="center">Type-safe HTTP APIs for Go. Declare a struct once and get validation, OpenAPI docs, a TypeScript client and Zod schemas.</p>
</td>
<td width="33%" valign="top">
<a href="https://github.com/raven-clown/idpforge"><img src="./project-images/covers/idpforge.png" alt="IdpForge"></a>
<h3 align="center"><a href="https://github.com/raven-clown/idpforge">IdpForge</a></h3>
<p align="center">Self-hosted SSO and identity platform in Go. OIDC, RBAC, realtime fan-out, and cluster-safe background jobs.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<a href="./experience/delta-electronics"><img src="./project-images/covers/ai-opensearch.png" alt="AI OpenSearch"></a>
<h3 align="center"><a href="./experience/delta-electronics">AI OpenSearch</a></h3>
<p align="center">Log analytics platform for the factory. Go collectors and workers, OpenSearch indices, and Teams alerts while production is still affected.</p>
</td>
<td width="33%" valign="top">
<a href="./experience/delta-electronics"><img src="./project-images/covers/central-kafka-platform.png" alt="Central Kafka Platform"></a>
<h3 align="center"><a href="./experience/delta-electronics">Central Kafka Platform</a></h3>
<p align="center">Shared Kafka and NiFi platforms other IT teams plug into, with CDC from the MES, Spark streaming jobs and Grafana monitoring.</p>
</td>
<td width="33%" valign="top">
<a href="./experience/fivem-redm-platform"><img src="./project-images/covers/fivem-multiplayer-platform.png" alt="FiveM Multiplayer Platform"></a>
<h3 align="center"><a href="./experience/fivem-redm-platform">FiveM Platform</a></h3>
<p align="center">A multiplayer roleplay platform I founded and ran for five years: game server, Lua systems, web shop and infrastructure.</p>
</td>
</tr>
</table>

<p align="center">
More: <a href="./experience/side-projects">side projects</a> · <a href="./experience/delta-electronics">work at Delta</a> · <a href="./experience/architecture.md">architecture diagrams</a> · <a href="./experience/academic-projects">academic projects</a>
</p>

<br>

## Open Source

Fixes and features sent upstream to projects I use. Merged work first, then what is still in review.

| Project | Contribution | Status |
|---|---|---|
| [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) | [#1042](https://github.com/mark3labs/mcp-go/pull/1042) streamable HTTP client reports the real HTTP status when a gateway answers with a JSON error body | Merged |
| [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) | [#1100](https://github.com/mark3labs/mcp-go/pull/1100) deterministic SSE custom client timeout test | Merged |
| [thebanri/limoni](https://github.com/thebanri/limoni) | [#110](https://github.com/thebanri/limoni/pull/110) accessibility tests for redaction, announcements and mode switches | Merged |
| [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) | [#1146](https://github.com/mark3labs/mcp-go/pull/1146) empty tool annotations omitted, `WithoutDefaultAnnotations` · [#1147](https://github.com/mark3labs/mcp-go/pull/1147) tool schema dialects and definitions preserved · [#1148](https://github.com/mark3labs/mcp-go/pull/1148) structured-only tool results · [#1149](https://github.com/mark3labs/mcp-go/pull/1149) notification params decoding · [#1150](https://github.com/mark3labs/mcp-go/pull/1150) replay callbacks get their own event copy · [#1151](https://github.com/mark3labs/mcp-go/pull/1151) empty PKCE challenge refused · [#1154](https://github.com/mark3labs/mcp-go/pull/1154) requests no longer re-parsed on the way to the handler · [#1155](https://github.com/mark3labs/mcp-go/pull/1155) listener retry interval | In review |
| [prometheus/blackbox_exporter](https://github.com/prometheus/blackbox_exporter) | [#1695](https://github.com/prometheus/blackbox_exporter/pull/1695) configurable EDNS0 UDP payload size · [#1696](https://github.com/prometheus/blackbox_exporter/pull/1696) configured IP protocol kept on cross-host redirects | In review |
| [grafana/mcp-grafana](https://github.com/grafana/mcp-grafana) | [#1455](https://github.com/grafana/mcp-grafana/pull/1455) TLS and SOCKS5 transport reused across tool calls | In review |
| [qax-os/excelize](https://github.com/qax-os/excelize) | [#2476](https://github.com/qax-os/excelize/pull/2476) set row height without `customHeight` | In review |
| [getgauge/gauge](https://github.com/getgauge/gauge) | [#2913](https://github.com/getgauge/gauge/pull/2913) escape backslash and pipe in table cells · [#2915](https://github.com/getgauge/gauge/pull/2915) keep multiline string arguments when formatting | In review |
| [open-policy-agent/regal](https://github.com/open-policy-agent/regal) | [#2150](https://github.com/open-policy-agent/regal/pull/2150) LSP path handling on Windows | In review |
| [clastix/kamaji](https://github.com/clastix/kamaji) | [#1403](https://github.com/clastix/kamaji/pull/1403) migration jobs scheduled from the DataStore | In review |

Also in review: [editorconfig-core-go #249](https://github.com/editorconfig/editorconfig-core-go/pull/249), [gograph #300](https://github.com/hmdsefi/gograph/pull/300), [mc-router #651](https://github.com/itzg/mc-router/pull/651), [sesh #526](https://github.com/joshmedeski/sesh/pull/526), [donetick #949](https://github.com/donetick/donetick/pull/949), [warewulf #2342](https://github.com/warewulf/warewulf/pull/2342), [doc-scraper #132](https://github.com/Sriram-PR/doc-scraper/pull/132).

<br>

## GitHub Stats

<div align="center">

<a href="https://github.com/raven-clown"><img src="https://github-profile-summary-cards.vercel.app/api/cards/profile-details?username=raven-clown&theme=transparent" alt="Profile details" width="100%"></a>

<table>
<tr>
<td width="33%"><img src="https://github-profile-summary-cards.vercel.app/api/cards/stats?username=raven-clown&theme=transparent" alt="Stats" width="100%"></td>
<td width="33%"><img src="https://github-profile-summary-cards.vercel.app/api/cards/most-commit-language?username=raven-clown&theme=transparent" alt="Top languages by commit" width="100%"></td>
<td width="33%"><img src="https://github-profile-summary-cards.vercel.app/api/cards/repos-per-language?username=raven-clown&theme=transparent" alt="Top languages by repo" width="100%"></td>
</tr>
</table>

<a href="https://github.com/raven-clown"><img src="https://streak-stats.demolab.com?user=raven-clown&hide_border=true&theme=transparent" alt="Contribution streak"></a>

</div>

<br>

## This Repository

| Folder | What's inside |
|---|---|
| [`experience/`](./experience) | The portfolio: what I built, where, and what came out of it |
| [`brain/`](./brain) | The knowledge base: how each skill area works, in detail |
| [`about/`](./about) | How I work and background outside the job title |

**Education:** B.Sc. Digital Business Technology, Southeast Bangkok University, 2025 (First Class Honors, Gold Medal)

<br>

<div align="center">

Open to interesting backend, platform and data engineering work. Reach me at [ekdanai.kk@gmail.com](mailto:ekdanai.kk@gmail.com).

</div>
