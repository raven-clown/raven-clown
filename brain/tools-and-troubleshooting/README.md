# Tools & Troubleshooting

How I approach environment and tooling problems — the work that has to happen before anything else in this repo is possible to ship.

## Debugging Installer/Environment Failures on Locked-Down Machines

- On a corporate, domain-joined Windows machine, an installer failing silently is often a Group Policy Object (GPO) blocking something specific — like packaged-app (MSIX) service installation — rather than the installer itself being broken. The fix is finding which policy is blocking it, not reinstalling repeatedly.
- A tool failing because an environment variable like `%USERPROFILE%` shows up literally in a path (instead of being expanded) usually means the tool was configured or invoked in a context where the shell isn't doing variable expansion the way it would interactively — worth checking exactly how and where the value was set.

## Tracing Network Issues Past the Obvious Cause

- A connection that "randomly" drops or retries isn't always the application's fault — MTU mismatches and packet fragmentation on the network path can produce exactly that symptom, and they're invisible unless you're specifically looking at the network layer instead of the application logs.
- Recovering a machine after cutting off your own remote access (for example, disabling a network adapter mid-session over RDP) usually means falling back to host-level access — console access, a hypervisor's remote console, or physical access — rather than anything network-dependent.

## Git History Surgery

- `git filter-repo` rewrites history rather than just adding a new commit on top — the right tool when something needs to be actually removed from every commit that touched it, not just reverted (a leaked secret, for example, since a revert still leaves the secret readable in the old commit).
- The same tool works for cleaning up unwanted metadata across history, like stripping co-author trailers that got added by a tool and shouldn't be part of the permanent record.
- Because it rewrites commit hashes, anyone with an existing clone needs to re-clone (or hard-reset to the new history) afterward — it's not a change that merges quietly.

## Why This Folder Exists

Most portfolios only show the finished feature. This folder is the other half — the environment and tooling debugging that has to happen before any of the other folders' work is even possible to ship.
