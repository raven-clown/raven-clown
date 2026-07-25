# FiveM & RedM Platform

**Founder & Full-Stack Developer** — Independent project · 2021 – 2025

## What it was

A multiplayer roleplay platform (highclass-roleplay.com / `sv_highclass`) built and operated on FiveM, later extended to RedM. Five years of ownership, end to end — server, backend systems, database, and interface design, none of it handed to me by a framework.

## What I built

**Core systems.** `raven_core`, the platform's economy and identity system (energy, SSNN/phone-number generation, batched database writes to hold up under concurrent players); `raven_pedscale`; `raven_advanceambulance`, an EMS system with its own knockdown/crawl state machine; `myopt`, a custom 7-module optimizer I wrote after reverse-engineering a Luraph-obfuscated third-party module to understand what it was actually doing.

**Database.** Redesigned the platform's database from an ad-hoc schema into a structured, indexed design built for batched writes under real load.

**Interface.** Designed and built the HUD and inventory UI (Svelte and React), and later carried the same design language into UX/UI work on RedM.

**Infrastructure.** Ran the platform on a dedicated Windows server (i9, 64GB RAM, 10Gbps), with an nginx caching layer in front for resource delivery and SSL on the storefront subdomain. Later rebuilt the website/shop side on Next.js and Go, deployed on AWS and Kubernetes.

**Tooling.** Built Raven Whitelist, a Discord-based whitelist management bot (TypeScript, Bun, Redis, BullMQ, MySQL), shipped as compiled binaries with bilingual documentation.

## Full technical write-up

See [`brain/fivem-redm-development`](../../brain/fivem-redm-development).
