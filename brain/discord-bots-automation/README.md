# Discord Bots & Automation

How a Discord bot and its background automation actually work.

## Bot Architecture

- A Discord bot maintains a persistent WebSocket connection (the "gateway") to Discord, which is how it receives real-time events (a message sent, a reaction added, a member joining) instead of polling for updates.
- Slash commands are registered with Discord ahead of time and show up as a native part of the Discord UI, which is different from older-style bots that just watched for a text prefix like `!command` in message content.
- Permissions and gateway intents control exactly what a bot is allowed to see and do. A bot has to explicitly declare which events it needs (message content, member updates, etc.), which limits what it can access even before per-server role permissions come into play.

## Job Queues (Redis + BullMQ)

- A job queue decouples "the thing that needs to happen" from "when it actually happens." Instead of doing slow or rate-limited work directly inside a command handler, the handler pushes a job onto the queue and returns immediately.
- BullMQ, backed by Redis, handles the producer/consumer side of that: jobs get added to a queue, one or more workers pick them up and process them, and failed jobs can be retried automatically instead of just disappearing.
- This matters most for anything rate-limited or bursty. A whitelist check that has to hit an external API, for example, benefits from being queued and processed at a controlled rate instead of firing every request the moment it comes in.

## Bun Runtime

- A JavaScript/TypeScript runtime built as a faster alternative to Node. Quicker cold start, and a built-in bundler and test runner instead of needing separate tools bolted on.
- Compiling down to standalone binaries means a bot can be distributed to someone running a server without them needing Node or any dependencies installed first. They just run the binary.

## Backend Layer

- The same NestJS + Prisma pattern used for web backends applies directly to a bot's backend. Commands map to controller-like handlers, business logic lives in services, and Prisma handles the typed data layer underneath. See [backend-development](../backend-development) for the full breakdown.
