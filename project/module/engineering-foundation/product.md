# Module Product: engineering-foundation

## Scope

- Module: M0 engineering skeleton.
- User value: A reproducible, tested base for milestone-by-milestone CodeKB development.
- Out of scope: Functional project initialization, storage, indexing, querying, or MCP.

## Functional Design

| Function | Description | Status | Notes |
| --- | --- | --- | --- |
| CLI root | Help/version and stable execution boundary | implemented | Cobra; no placeholder commands |
| Config contract | Versioned defaults and value validation | implemented | Filesystem loading belongs to M1 |
| Diagnostics | JSON logger targeting caller-provided stderr | implemented | Standard library slog |
| Automation | Format, vet, test, and build locally/in CI | implemented | macOS CI; local arm64 validation |

## API / UI / Data Surface

| Surface | Direction | Parameters / Fields | Status |
| --- | --- | --- | --- |
| `codekb --help` | user → CLI | none | implemented |
| `codekb --version` | user → CLI | build version | implemented |
| `config.Config` | internal | schema/project/index/search/embedding | implemented |

## Rules And States

- CLI failures return non-zero and write errors to stderr.
- Embedding is disabled by default.
- Unimplemented milestone commands are not advertised.

## Dependencies

- Go standard library.
- `github.com/spf13/cobra`.

## Implementation Status

| Item | Status | Code / Doc Reference | Notes |
| --- | --- | --- | --- |
| M0 skeleton | implemented | `cmd/`, `internal/`, `README.md`, `Makefile` | Checks and smoke tests pass |
