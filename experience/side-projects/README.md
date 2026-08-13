# Side Projects

## LINE Task Bot

A group task-management bot for LINE, built to run on free-tier infrastructure end to end (NestJS, Prisma, PostgreSQL on Render/Supabase). Task creation, exclusive vs. open-claim task modes, role permissions, deadline notifications — shipped with 10 markdown docs covering the schema, MVP roadmap, and a monetization plan (roughly 49–99 THB per group per month) before a single user signed up.

## Excel Habit Tracker

A habit-tracking workbook generated with Python (`openpyxl`) instead of built by hand — monthly grids, 20 recurring tasks plus 5 ad-hoc slots a month, a GitHub-style contribution heatmap, and a yearly KPI summary.

## Snippet Manager

A code snippet manager built as two front-ends — a VS Code extension and an Electron desktop app — that read and write the same local `data.json` and stay in sync in real time via file watching (chokidar), with no server, no account, and no cloud involved. Started August 11, 2026 and under active development since; currently at v0.6.0, MIT-licensed, in an npm workspaces monorepo (`raven-clown/monorepo`) split into `@snippet/core` (storage, CRUD, and file-watch sync — no UI), `vscode-extension`, and `electron-app`. All storage, validation, and migration logic lives in `@snippet/core`; neither front-end touches the filesystem directly. Both apps read and write the same `data.json` in the OS's standard app-data directory (`%APPDATA%\snippet-manager` on Windows) and watch it with chokidar — no polling, no direct IPC between the two.

A snippet holds title, code, language, tags, category, pinned state, usage tracking, and a hidden-from-VS-Code flag. Categories are a real tree entity — id/parentId/order/pinned — supporting unlimited nesting, migrated automatically from the old flat-string schema. Export/import covers both full-store backups and single snippets, with merge or replace modes.

The desktop app is the fuller surface: a glass/blur "iOS 26" visual style across three themes (White/Black/Color), generated from one set of design tokens in the oklch color space — first prototyped in Claude Design, then converted into real CSS custom properties rather than an off-the-shelf UI library. A VS Code Explorer-style category tree with drag-and-drop reorder/reparent, inline rename, and inline code editing directly on the detail view. A from-scratch syntax highlighter. A custom frameless title bar. Full EN/TH i18n. A responsive layout across three breakpoints (compact/medium/wide) where the sidebar collapses into a drawer on narrow screens. The storage location itself is relocatable, with a pointer-file mechanism to keep sync safe across the move.

The VS Code extension is the lighter surface — a sidebar tree grouped by language with a pinned section floating to the top, auto-reveal of the active file's language group, click-to-copy/insert/open per setting, and filtering out snippets marked hidden-from-VS-Code on the desktop side.

Distributed via electron-builder (nsis-web) with GitHub Actions building and publishing installers to GitHub Releases on every tag push (`v*`), and in-app auto-update through electron-updater against the same release feed. Currently building a custom-branded installer to replace the default NSIS UI. Stack: TypeScript, Electron 43, React 18, Vite (electron-vite), chokidar, Webpack, VS Code Extension API.

## Full technical write-up

See [`brain/side-projects-architecture`](../../brain/side-projects-architecture) and [`brain/discord-bots-automation`](../../brain/discord-bots-automation).
