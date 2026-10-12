# Module Todo: engineering-foundation

## Current Focus

- M0 is complete; wait for user confirmation before M1.

## In Progress

| Task | Status | Owner/Context | Next Step |
| --- | --- | --- | --- |

## Planned

| Task | Status | Priority | Notes |
| --- | --- | --- | --- |
| Inject release version at build time | planned | low | Add with release workflow when needed |

## Blocked

| Task | Status | Blocker | Needed Decision |
| --- | --- | --- | --- |

## Implemented

| Date | Task | Verification | Notes |
| --- | --- | --- | --- |
| 2026-10-12 | Buildable CLI and foundation modules | `make check`; `go test -race ./...`; help/version/error smoke tests | M0 complete |

## Assumptions

- Module path is `github.com/yanglichao/codekb`, inferred from the workspace owner and project name.
- Go 1.25 is the minimum baseline because it is the installed stable toolchain in this development environment.
