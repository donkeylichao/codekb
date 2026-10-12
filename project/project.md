# Project Overview

## Summary

- Project: CodeKB
- Goal: Build an offline-first, project-local Go code knowledge base and read-only MCP server.
- Main modules: engineering foundation, project/storage, indexing, search/graph, MCP, incremental/semantic retrieval.

## Implemented

| Feature | Module | Status | Notes |
| --- | --- | --- | --- |
| Go/Cobra engineering skeleton | engineering-foundation | implemented | `make check` and race tests pass |

## Not Implemented

| Feature | Module | Status | Notes |
| --- | --- | --- | --- |
| Project initialization and SQLite | project-storage | planned | M1 |
| Go/document indexing | indexing | planned | M2 |
| Search and graph queries | search-graph | planned | M3 |
| MCP and Codex registration | mcp | planned | M4 |
| Semantic and incremental indexing | incremental-semantic | planned | M5 |

## Architecture Notes

- Single Go binary; project-local state; offline by default.
- CLI and future MCP transport share application services.
- Source is authoritative and indexes are disposable caches.

## Recent Updates

| Date | Module | Change | Verification |
| --- | --- | --- | --- |
| 2026-10-12 | engineering-foundation | Completed M0 skeleton | `make check`; `go test -race ./...`; CLI smoke tests |
