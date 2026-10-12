# Product Plan

## Product Direction

- Target users: Go developers and AI-assisted engineering workflows.
- Core value: Traceable symbol, call-graph, documentation, and search evidence without uploading source.
- Product boundaries: Go + Markdown/TXT first; no GUI, cloud model, distributed graph database, or automatic source modification.

## Feature Modules

| Module | Product Scope | Status | Notes |
| --- | --- | --- | --- |
| engineering-foundation | Buildable CLI, config contract, diagnostics, CI | implemented | M0 verified |
| project-storage | Init, validation, SQLite, status | planned | M1 |
| indexing | Go symbols/calls and document chunks | planned | M2 |
| search-graph | FTS and graph traversal | planned | M3 |
| mcp | Read-only stdio tools and safe Codex registration | planned | M4 |
| incremental-semantic | Updates and optional Ollama retrieval | planned | M5 |

## Status Values

- `planned`: planned but not started
- `in_progress`: currently being implemented
- `implemented`: implemented and verified or verification noted
- `blocked`: blocked by missing decision, dependency, or external condition
- `deferred`: intentionally postponed

## Cross-Module Requirements

| Requirement | Modules | Status | Notes |
| --- | --- | --- | --- |
| Project isolation and source evidence | all | planned | Enforced as modules are added |
| Local-only operation by default | all | planned | Config default implemented in M0 |
| macOS arm64 support | all | implemented | M0 built and tested on darwin/arm64 |

## Open Questions

- SQLite driver and FTS5 support will be selected and verified in M1.
- Current Codex project configuration behavior will be verified before M4.
