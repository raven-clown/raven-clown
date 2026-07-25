# Side Projects

## Punya — Super-App

Designed the full architecture for a multi-feature app: a social webboard plus "Task-Lock," an immutable productivity engine, across 8 user tiers with its own monetization, auth, and RBAC design. Originally scoped as a Go API gateway with NestJS microservices on GKE; once I priced out what that would actually cost to run, I re-scoped it down to a single Hetzner VPS running Docker Compose. 14+ markdown docs cover the design before any production code was written. Stack: Go, TypeScript, PostgreSQL/Drizzle, Redis Streams.

## Synapse — Flashcard App

A complete NestJS backend for a spaced-repetition flashcard app: JWT auth, the SM-2 scheduling algorithm, an EXP/level system, streak tracking, a friend system, Redis caching, Swagger docs, rate limiting, and refresh-token rotation. [github.com/AK47-DEVELOPER/synapse](https://github.com/AK47-DEVELOPER/synapse)

## LINE Task Bot

A group task-management bot for LINE, built to run on free-tier infrastructure end to end (NestJS, Prisma, PostgreSQL on Render/Supabase). Task creation, exclusive vs. open-claim task modes, role permissions, deadline notifications — shipped with 10 markdown docs covering the schema, MVP roadmap, and a monetization plan (roughly 49–99 THB per group per month) before a single user signed up.

## Excel Habit Tracker

A habit-tracking workbook generated with Python (`openpyxl`) instead of built by hand — monthly grids, 20 recurring tasks plus 5 ad-hoc slots a month, a GitHub-style contribution heatmap, and a yearly KPI summary.

## Full technical write-up

See [`brain/side-projects-architecture`](../../brain/side-projects-architecture) and [`brain/discord-bots-automation`](../../brain/discord-bots-automation).
