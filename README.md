<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./assets/header-dark.svg">
  <img src="./assets/header-light.svg" alt="Ekdanai Kummee, Application Engineer. DevOps, data platforms and full-stack web." width="100%">
</picture>

<div align="center">

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./assets/ravenfetch-dark.svg">
  <img src="./assets/ravenfetch-light.svg" alt="ravenfetch: my profile as a terminal program. The raven from my avatar, drawn in ASCII with a flickering lantern, next to who I am, what I build and where to reach me." width="100%">
</picture>

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
- In my own time I build the mothership stack below, open source Go infrastructure for running agents, and send fixes upstream to the tools I use.

<br>

## Building Now: the Mothership Stack

Four Go projects for running agents in production. Each one works alone, and they fit together through [one shared contract](https://github.com/raven-clown/mothership). Everything ships in the free build: SSO, roles, audit logs, clustering.

<table>
<tr>
<td width="25%" align="center" valign="top">
<a href="https://github.com/raven-clown/airlock"><img src="./assets/logos/airlock.svg" width="72" alt="airlock logo"></a>
<h3><a href="https://github.com/raven-clown/airlock">airlock</a></h3>
<a href="https://github.com/raven-clown/airlock/releases"><img src="https://img.shields.io/github/v/release/raven-clown/airlock?style=flat-square&color=6366F1&label=release" alt="airlock release"></a>
<p>MCP gateway. One front door for many MCP servers, with login, permissions, an audit log and rate limits.</p>
</td>
<td width="25%" align="center" valign="top">
<a href="https://github.com/raven-clown/rosetta"><img src="./assets/logos/rosetta.svg" width="72" alt="rosetta logo"></a>
<h3><a href="https://github.com/raven-clown/rosetta">rosetta</a></h3>
<a href="https://github.com/raven-clown/rosetta/releases"><img src="https://img.shields.io/github/v/release/raven-clown/rosetta?style=flat-square&color=F97316&label=release" alt="rosetta release"></a>
<p>Agent protocol bridge. Write an agent once and serve it over MCP, A2A, ACP and AG-UI.</p>
</td>
<td width="25%" align="center" valign="top">
<a href="https://github.com/raven-clown/transponder"><img src="./assets/logos/transponder.svg" width="72" alt="transponder logo"></a>
<h3><a href="https://github.com/raven-clown/transponder">transponder</a></h3>
<a href="https://github.com/raven-clown/transponder/releases"><img src="https://img.shields.io/github/v/release/raven-clown/transponder?style=flat-square&color=A855F7&label=release" alt="transponder release"></a>
<p>LLM gateway. One API in front of Anthropic, OpenAI, Gemini, Ollama and more, with fallback, budgets and token limits.</p>
</td>
<td width="25%" align="center" valign="top">
<a href="https://github.com/raven-clown/logbook"><img src="./assets/logos/logbook.svg" width="72" alt="logbook logo"></a>
<h3><a href="https://github.com/raven-clown/logbook">logbook</a></h3>
<a href="https://github.com/raven-clown/logbook/releases"><img src="https://img.shields.io/github/v/release/raven-clown/logbook?style=flat-square&color=10B981&label=release" alt="logbook release"></a>
<p>Agent memory. A shared memory any agent can read and write, served as an MCP server.</p>
</td>
</tr>
</table>

<br>

## Tech Stack

<div align="center">

<table>
<tr>
<td align="right"><b>Languages</b></td>
<td><a href="https://skillicons.dev"><img src="https://skillicons.dev/icons?i=go,cs,ts,js,py,java,php,lua,html,css,bash,powershell&perline=12" alt="Languages"></a></td>
</tr>
<tr>
<td align="right"><b>Frontend</b></td>
<td><a href="https://skillicons.dev"><img src="https://skillicons.dev/icons?i=react,nextjs,nuxtjs,vue,svelte,jquery,tailwind,bootstrap,vite,electron,figma,ps&perline=12" alt="Frontend"></a></td>
</tr>
<tr>
<td align="right"><b>Backend</b></td>
<td><a href="https://skillicons.dev"><img src="https://skillicons.dev/icons?i=nodejs,bun,dotnet,nestjs,express,fastapi,spring,laravel,prisma,npm,vitest,ai&perline=12" alt="Backend"></a></td>
</tr>
<tr>
<td align="right"><b>Data and messaging</b></td>
<td><a href="https://skillicons.dev"><img src="https://skillicons.dev/icons?i=kafka,rabbitmq,redis,postgres,mysql,sqlite,mongodb,supabase,elasticsearch&perline=12" alt="Data and messaging"></a></td>
</tr>
<tr>
<td align="right"><b>Infrastructure</b></td>
<td><a href="https://skillicons.dev"><img src="https://skillicons.dev/icons?i=docker,kubernetes,nginx,linux,redhat,windows,aws,vercel,gitlab,githubactions,grafana,prometheus&perline=12" alt="Infrastructure"></a></td>
</tr>
<tr>
<td align="right"><b>Tools</b></td>
<td><a href="https://skillicons.dev"><img src="https://skillicons.dev/icons?i=git,github,vscode,visualstudio,notion,discord,arduino&perline=12" alt="Tools"></a></td>
</tr>
</table>

![Apache NiFi](https://img.shields.io/badge/Apache_NiFi-728E9B?style=flat-square&logo=apachenifi&logoColor=white)
![Apache Airflow](https://img.shields.io/badge/Apache_Airflow-017CEE?style=flat-square&logo=apacheairflow&logoColor=white)
![Apache Spark](https://img.shields.io/badge/Apache_Spark-E25A1C?style=flat-square&logo=apachespark&logoColor=white)
![OpenSearch](https://img.shields.io/badge/OpenSearch-005EB8?style=flat-square&logo=opensearch&logoColor=white)
![NATS](https://img.shields.io/badge/NATS_JetStream-27AAE1?style=flat-square&logo=natsdotio&logoColor=white)
![AKHQ](https://img.shields.io/badge/AKHQ-00ACC1?style=flat-square)
![Oracle](https://img.shields.io/badge/Oracle-F80000?style=flat-square&logo=oracle&logoColor=white)
![SQL Server](https://img.shields.io/badge/SQL_Server-CC2927?style=flat-square)
![MinIO](https://img.shields.io/badge/MinIO-C72E49?style=flat-square&logo=minio&logoColor=white)
![Kong](https://img.shields.io/badge/Kong-003459?style=flat-square&logo=kong&logoColor=white)
![Harbor](https://img.shields.io/badge/Harbor-60B932?style=flat-square)
![Rancher](https://img.shields.io/badge/Rancher-0075A8?style=flat-square&logo=rancher&logoColor=white)
![Docker Swarm](https://img.shields.io/badge/Docker_Swarm-2496ED?style=flat-square&logo=docker&logoColor=white)
![IIS](https://img.shields.io/badge/IIS-5E5E5E?style=flat-square&logo=windows&logoColor=white)
![Windows Server](https://img.shields.io/badge/Windows_Server-0078D4?style=flat-square&logo=windows&logoColor=white)
![RHEL](https://img.shields.io/badge/RHEL-EE0000?style=flat-square&logo=redhat&logoColor=white)
![Playwright](https://img.shields.io/badge/Playwright-2EAD33?style=flat-square)
![MCP](https://img.shields.io/badge/MCP-000000?style=flat-square&logo=modelcontextprotocol&logoColor=white)
![Ollama](https://img.shields.io/badge/Ollama-000000?style=flat-square&logo=ollama&logoColor=white)
![LINE API](https://img.shields.io/badge/LINE_API-06C755?style=flat-square&logo=line&logoColor=white)
![Cisco](https://img.shields.io/badge/Cisco-1BA0D7?style=flat-square&logo=cisco&logoColor=white)
![Power BI](https://img.shields.io/badge/Power_BI-F2C811?style=flat-square&logo=powerbi&logoColor=black)
![Machine Learning](https://img.shields.io/badge/Machine_Learning-FF6F00?style=flat-square)
![Adobe Illustrator](https://img.shields.io/badge/Illustrator-FF9A00?style=flat-square)

</div>

<br>

## Featured Projects

<table>
<tr>
<td width="33%" valign="top">
<a href="https://github.com/raven-clown/ark"><img src="./project-images/covers/ark-kafkatohttp-bridge.png" alt="ARK"></a>
<h3 align="center"><a href="https://github.com/raven-clown/ark">ARK</a></h3>
<p align="center"><img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"> <img src="https://img.shields.io/badge/Kafka-231F20?style=flat-square&logo=apachekafka&logoColor=white" alt="Kafka"> <img src="https://img.shields.io/badge/React-20232A?style=flat-square&logo=react&logoColor=61DAFB" alt="React"></p>
<p align="center">Kafka-to-HTTP bridge. Retries, dead letters, flows, cluster mode and a live console. One binary, one YAML file.</p>
</td>
<td width="33%" valign="top">
<a href="https://github.com/raven-clown/tack"><img src="./project-images/covers/tack-typesafe-http-apis-for-go.png" alt="tack"></a>
<h3 align="center"><a href="https://github.com/raven-clown/tack">tack</a></h3>
<p align="center"><img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"> <img src="https://img.shields.io/badge/OpenAPI_3.1-6BA539?style=flat-square&logo=openapiinitiative&logoColor=white" alt="OpenAPI"> <img src="https://img.shields.io/badge/TypeScript-3178C6?style=flat-square&logo=typescript&logoColor=white" alt="TypeScript"></p>
<p align="center">Type-safe HTTP APIs for Go. Declare a struct once and get validation, OpenAPI docs, a TypeScript client and Zod schemas.</p>
</td>
<td width="33%" valign="top">
<a href="https://github.com/raven-clown/idpforge"><img src="./project-images/covers/idpforge.png" alt="IdpForge"></a>
<h3 align="center"><a href="https://github.com/raven-clown/idpforge">IdpForge</a></h3>
<p align="center"><img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"> <img src="https://img.shields.io/badge/OIDC-F78C40?style=flat-square&logo=openid&logoColor=white" alt="OIDC"> <img src="https://img.shields.io/badge/Next.js-000000?style=flat-square&logo=nextdotjs&logoColor=white" alt="Next.js"></p>
<p align="center">Self-hosted SSO and identity platform. OIDC, RBAC, realtime fan-out, and cluster-safe background jobs.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<a href="./experience/delta-electronics"><img src="./project-images/covers/ai-opensearch.png" alt="AI OpenSearch"></a>
<h3 align="center"><a href="./experience/delta-electronics">AI OpenSearch</a></h3>
<p align="center"><img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go"> <img src="https://img.shields.io/badge/OpenSearch-005EB8?style=flat-square&logo=opensearch&logoColor=white" alt="OpenSearch"> <img src="https://img.shields.io/badge/FastAPI-009688?style=flat-square&logo=fastapi&logoColor=white" alt="FastAPI"></p>
<p align="center">Log analytics for the factory. Go collectors and workers, OpenSearch indices, and Teams alerts while production is still affected.</p>
</td>
<td width="33%" valign="top">
<a href="./experience/delta-electronics"><img src="./project-images/covers/central-kafka-platform.png" alt="Central Kafka Platform"></a>
<h3 align="center"><a href="./experience/delta-electronics">Central Kafka Platform</a></h3>
<p align="center"><img src="https://img.shields.io/badge/Kafka-231F20?style=flat-square&logo=apachekafka&logoColor=white" alt="Kafka"> <img src="https://img.shields.io/badge/NiFi-728E9B?style=flat-square&logo=apachenifi&logoColor=white" alt="NiFi"> <img src="https://img.shields.io/badge/Spark-E25A1C?style=flat-square&logo=apachespark&logoColor=white" alt="Spark"></p>
<p align="center">Shared Kafka and NiFi platforms other IT teams plug into, with CDC from the MES, Spark streaming jobs and Grafana monitoring.</p>
</td>
<td width="33%" valign="top">
<a href="./experience/fivem-redm-platform"><img src="./project-images/covers/fivem-multiplayer-platform.png" alt="FiveM Multiplayer Platform"></a>
<h3 align="center"><a href="./experience/fivem-redm-platform">FiveM Platform</a></h3>
<p align="center"><img src="https://img.shields.io/badge/Lua-2C2D72?style=flat-square&logo=lua&logoColor=white" alt="Lua"> <img src="https://img.shields.io/badge/MySQL-4479A1?style=flat-square&logo=mysql&logoColor=white" alt="MySQL"> <img src="https://img.shields.io/badge/Svelte-FF3E00?style=flat-square&logo=svelte&logoColor=white" alt="Svelte"></p>
<p align="center">A multiplayer roleplay platform I founded and ran for five years: game server, Lua systems, web shop and infrastructure.</p>
</td>
</tr>
</table>

<p align="center">
More: <a href="./experience/side-projects">side projects</a> · <a href="./experience/delta-electronics">work at Delta</a> · <a href="./experience/architecture.md">architecture diagrams</a> · <a href="./experience/academic-projects">academic projects</a>
</p>

<br>

## Open Source

20 pull requests merged upstream, in 11 projects I use.

<table>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center"><img src="https://img.shields.io/badge/merged-10_PRs-8957e5?style=flat-square&logo=github&logoColor=white" alt="10 merged PRs"> <img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars"></p>
<p align="center">The Go SDK for MCP. Transport errors that show the real status, stdio and PKCE edge cases, safe event replay, race fixes, and a request path that no longer parses the same JSON several times.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/grafana/mcp-grafana">grafana/mcp-grafana</a></h3>
<p align="center"><a href="https://github.com/grafana/mcp-grafana/pull/1455"><img src="https://img.shields.io/badge/merged-%231455-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1455"></a> <img src="https://img.shields.io/github/stars/grafana/mcp-grafana?style=flat-square&color=555" alt="stars"></p>
<p align="center">Tool calls with custom TLS or a SOCKS5 proxy share one transport. 300 calls open 1 connection instead of 300, and each call is about 25 times faster.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/joshmedeski/sesh">joshmedeski/sesh</a></h3>
<p align="center"><a href="https://github.com/joshmedeski/sesh/pull/526"><img src="https://img.shields.io/badge/merged-%23526-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 526"></a> <img src="https://img.shields.io/github/stars/joshmedeski/sesh?style=flat-square&color=555" alt="stars"></p>
<p align="center">Every lookup uses the configured <code>tmux_command</code>, so another multiplexer such as psmux works for all commands.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/clastix/kamaji">clastix/kamaji</a></h3>
<p align="center"><a href="https://github.com/clastix/kamaji/pull/1403"><img src="https://img.shields.io/badge/merged-%231403-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1403"></a> <img src="https://img.shields.io/github/stars/clastix/kamaji?style=flat-square&color=555" alt="stars"></p>
<p align="center">DataStore migrations take node selectors, tolerations, affinity and a security context, so the migration job runs on tainted and hardened nodes.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/itzg/mc-router">itzg/mc-router</a></h3>
<p align="center"><a href="https://github.com/itzg/mc-router/pull/651"><img src="https://img.shields.io/badge/merged-%23651-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 651"></a> <img src="https://img.shields.io/github/stars/itzg/mc-router?style=flat-square&color=555" alt="stars"></p>
<p align="center">A <code>-log-file</code> option writes logs to a file as well, so fail2ban can read them from an image that has no shell.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/warewulf/warewulf">warewulf/warewulf</a></h3>
<p align="center"><a href="https://github.com/warewulf/warewulf/pull/2342"><img src="https://img.shields.io/badge/merged-%232342-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 2342"></a> <img src="https://img.shields.io/github/stars/warewulf/warewulf?style=flat-square&color=555" alt="stars"></p>
<p align="center">The firewalld service opens the TLS port too, so nodes with <code>tls: true</code> get their runtime overlay instead of a dropped connection.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/open-policy-agent/regal">open-policy-agent/regal</a></h3>
<p align="center"><a href="https://github.com/open-policy-agent/regal/pull/2150"><img src="https://img.shields.io/badge/merged-%232150-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 2150"></a> <img src="https://img.shields.io/github/stars/open-policy-agent/regal?style=flat-square&color=555" alt="stars"></p>
<p align="center">The Rego language server handles Windows paths, which fixed <code>go test</code> failures in four packages on Windows.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/hmdsefi/gograph">hmdsefi/gograph</a></h3>
<p align="center"><a href="https://github.com/hmdsefi/gograph/pull/300"><img src="https://img.shields.io/badge/merged-%23300-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 300"></a> <a href="https://github.com/hmdsefi/gograph/pull/302"><img src="https://img.shields.io/badge/merged-%23302-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 302"></a> <img src="https://img.shields.io/github/stars/hmdsefi/gograph?style=flat-square&color=555" alt="stars"></p>
<p align="center">Two graph algorithms: <code>NegativeCycle</code> returns a negative-weight cycle, and <code>Bidirectional</code> finds the cheapest path by searching from both ends.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/Sriram-PR/doc-scraper">Sriram-PR/doc-scraper</a> · <a href="https://github.com/thebanri/limoni">thebanri/limoni</a></h3>
<p align="center"><a href="https://github.com/Sriram-PR/doc-scraper/pull/132"><img src="https://img.shields.io/badge/merged-%23132-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 132"></a> <a href="https://github.com/thebanri/limoni/pull/110"><img src="https://img.shields.io/badge/merged-%23110-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 110"></a></p>
<p align="center">Pooled Markdown converters took crawls from 559 to about 905 pages per second, and accessibility tests raised a terminal UI engine's coverage from 65% to 91%.</p>
</td>
</tr>
</table>

<details>
<summary><b>Every pull request, one by one</b></summary>

<br>

| Project | PR | What changed |
|---|---|---|
| mark3labs/mcp-go | [#1172](https://github.com/mark3labs/mcp-go/pull/1172) | The last stdio message is kept when the server exits without a trailing newline, instead of failing with `transport closed`. |
| clastix/kamaji | [#1403](https://github.com/clastix/kamaji/pull/1403) | Scheduling fields and a security context for DataStore migration jobs. |
| itzg/mc-router | [#651](https://github.com/itzg/mc-router/pull/651) | `-log-file` appends logs to a file as well as stdout. |
| warewulf/warewulf | [#2342](https://github.com/warewulf/warewulf/pull/2342) | Opens the TLS port 9874 in the firewalld service. |
| open-policy-agent/regal | [#2150](https://github.com/open-policy-agent/regal/pull/2150) | Fixes path handling on Windows in the language server. |
| hmdsefi/gograph | [#300](https://github.com/hmdsefi/gograph/pull/300) | `path.NegativeCycle` returns a negative-weight cycle. |
| hmdsefi/gograph | [#302](https://github.com/hmdsefi/gograph/pull/302) | `path.Bidirectional` for the cheapest path between two vertices. |

<table>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1042"><img src="https://img.shields.io/badge/Merged-%231042-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1042"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">The streamable HTTP client now reports the real HTTP status when a gateway or framework answers with a JSON error body, instead of failing with an unrelated parse error.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1100"><img src="https://img.shields.io/badge/Merged-%231100-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1100"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">A flaky SSE transport test raced a 1 ns client timeout against a mock server and panicked on a nil error about half the time on Windows. It is deterministic now.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1146"><img src="https://img.shields.io/badge/Merged-%231146-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1146"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Tools no longer send an empty <code>annotations</code> object, and the new <code>WithoutDefaultAnnotations()</code> drops the default hints, so clients can tell an annotated tool from one without hints.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1148"><img src="https://img.shields.io/badge/Merged-%231148-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1148"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Clients accept a tool result that only has <code>structuredContent</code>, instead of failing with <code>content is missing</code>.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1149"><img src="https://img.shields.io/badge/Merged-%231149-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1149"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Decoding a notification into a reused value no longer keeps <code>_meta</code> and fields left over from the previous notification.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1150"><img src="https://img.shields.io/badge/Merged-%231150-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1150"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Each replay callback gets its own copy of a stored event, so a callback that changes a message can no longer corrupt what later replays send.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1151"><img src="https://img.shields.io/badge/Merged-%231151-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1151"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">OAuth with PKCE enabled now refuses an empty code challenge or verifier, instead of quietly finishing the flow without PKCE.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1154"><img src="https://img.shields.io/badge/Merged-%231154-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1154"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Requests are no longer parsed several times on the way to the handler. Profiling showed those extra JSON reads took most of the time and memory on large tool calls.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/mark3labs/mcp-go">mark3labs/mcp-go</a></h3>
<p align="center">
<a href="https://github.com/mark3labs/mcp-go/pull/1155"><img src="https://img.shields.io/badge/Merged-%231155-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1155"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/mark3labs/mcp-go?style=flat-square&color=555" alt="stars">
</p>
<p align="center">The client listener keeps the retry interval it was created with instead of reading a package variable, which fixes an intermittent <code>go test -race</code> failure.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/grafana/mcp-grafana">grafana/mcp-grafana</a></h3>
<p align="center">
<a href="https://github.com/grafana/mcp-grafana/pull/1455"><img src="https://img.shields.io/badge/Merged-%231455-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 1455"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/grafana/mcp-grafana?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Tool calls with a custom TLS config or a SOCKS5 proxy share one transport instead of building a new one each time. 300 calls open 1 connection instead of 300, and each call is about 25 times faster.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/Sriram-PR/doc-scraper">Sriram-PR/doc-scraper</a></h3>
<p align="center">
<a href="https://github.com/Sriram-PR/doc-scraper/pull/132"><img src="https://img.shields.io/badge/Merged-%23132-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 132"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/Sriram-PR/doc-scraper?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Markdown converters are pooled, so crawls keep scaling past 8 workers. The maintainer measured about 905 pages per second at 32 workers, against 559 before.</p>
</td>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/joshmedeski/sesh">joshmedeski/sesh</a></h3>
<p align="center">
<a href="https://github.com/joshmedeski/sesh/pull/526"><img src="https://img.shields.io/badge/Merged-%23526-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 526"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/joshmedeski/sesh?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Window, pane and session lookups use the configured <code>tmux_command</code>, so another multiplexer such as psmux works for every command, not only some.</p>
</td>
</tr>
<tr>
<td width="33%" valign="top">
<h3 align="center"><a href="https://github.com/thebanri/limoni">thebanri/limoni</a></h3>
<p align="center">
<a href="https://github.com/thebanri/limoni/pull/110"><img src="https://img.shields.io/badge/Merged-%23110-8957e5?style=flat-square&logo=github&logoColor=white" alt="Merged PR 110"></a>
<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go">
<img src="https://img.shields.io/github/stars/thebanri/limoni?style=flat-square&color=555" alt="stars">
</p>
<p align="center">Tests for the accessibility package of a Go terminal UI engine: value redaction, screen reader announcements on three platforms, and mode switches. Coverage went from 65% to 91%.</p>
</td>
</tr>
</table>

</details>

<br>

## Activity

<div align="center">

<img src="./generated/contrib-3d.svg" alt="3D contribution graph" width="100%">

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./generated/snake-dark.svg">
  <img src="./generated/snake-light.svg" alt="Contribution snake" width="100%">
</picture>

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

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="./assets/footer-dark.svg">
  <img src="./assets/footer-light.svg" alt="" width="100%">
</picture>
