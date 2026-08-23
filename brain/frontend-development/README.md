# Frontend Development

How the frameworks and design work actually function, point by point.

## React

- Built around components: small, composable functions that return UI and can hold their own state via hooks (`useState`, `useEffect`, and friends).
- Uses a virtual DOM. React builds an in-memory representation of the UI, diffs it against the previous version on every state change, and only updates the real DOM nodes that actually changed.
- Data flows one way, parent to child, via props; state that multiple components need gets lifted up to their common parent (or into a shared store) rather than passed sideways.

## Next.js

- Extends React with file-based routing. A file's path under the `pages` or `app` directory determines its URL, no separate router configuration needed.
- Supports multiple rendering modes for the same app: server-side rendering per request, static generation at build time, or incremental static regeneration that rebuilds a page in the background on a schedule.
- API routes let a Next.js app expose its own backend endpoints alongside the frontend, which is useful for small services that don't need a fully separate backend.

## Vue / Nuxt.js

- Vue's reactivity system tracks which parts of the UI depend on which piece of state, so an update to that state only re-renders the specific components that actually use it.
- Templates are closer to plain HTML than JSX is, with directives (`v-if`, `v-for`) handling conditionals and loops declaratively in the markup itself.
- Nuxt sits on top of Vue the way Next.js sits on top of React: file-based routing, SSR, and build tooling, pre-configured instead of assembled by hand.

## Svelte

- The biggest structural difference from React/Vue: there's no virtual DOM. Svelte is a compiler; it turns component code into direct, imperative DOM-update instructions at build time, so there's no diffing step at runtime.
- That makes it a strong fit for UI that has to update very frequently without dropping frames. A real-time status display or a HUD is the kind of case where skipping the diffing overhead actually matters.
- Reactivity is built into the language syntax itself (`$:` labeled statements re-run automatically when their dependencies change) rather than requiring an explicit hook.

## UX/UI Design Practice

- Design work happens in Figma: frames for each screen/state, a shared component library so a button or input looks the same everywhere it's used, and a consistent set of spacing/type tokens instead of eyeballing values per screen.
- Consistency first. Spacing, color, and interaction patterns should repeat predictably so a user doesn't have to relearn the interface screen to screen.
- Every action needs visible feedback. A click, a toggle, a state change should be confirmed on screen immediately, especially in an interface layered on top of something else happening in real time underneath it.
- Design for the failure/edge case, not just the happy path. What the screen looks like when data hasn't loaded yet, or when an action fails, gets designed deliberately rather than left as a blank state.

## Electron

- A desktop app that's really two processes: a Node-capable main process (windows, native menus, filesystem, auto-update) and one or more renderer processes that run web content (React, Vite) with Node access disabled by default and a preload script bridging the two through a controlled IPC surface.
- A frameless/custom title bar means giving up the OS chrome and rebuilding window controls (minimize/maximize/close, drag regions) by hand in the renderer. More design control, more edge cases to handle across platforms.
- Distribution goes through electron-builder to produce platform installers, and electron-updater to check a release feed (GitHub Releases works directly) and apply updates in-app instead of asking the user to redownload manually.
