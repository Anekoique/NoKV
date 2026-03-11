# Project Layout

This repository is organized by runtime domain rather than by a flat package list.

## Top-level domains

- `engine/`: embeddable NoKV API, write/read pipeline, iterators, stats, and options
- `storage/`: local persistence and indexing
- `cluster/`: distributed data plane
- `pd/`: PD-lite routing, TSO, and persistence
- `observability/`: metrics and hot-key tracking
- `tools/`: scripts and benchmark harness
- `tests/`: black-box and repo-layout tests
- `cmd/`: shipping binaries only

## Layout rules

- The repository root contains no Go source files.
- Embedded library consumers import `github.com/feichai0017/NoKV/engine`.
- White-box tests stay next to implementation when they need unexported access.
- Black-box and repository-shape checks live under `tests/`.
- Non-shipping developer utilities live under `tools/`.

## Suggested reading order

1. `engine/`
2. `storage/`
3. `cluster/`
4. `pd/`
5. `observability/`
6. `cmd/`
7. `tools/`

The README files in each top-level domain give a short map before you dive into package internals. For the path-by-path cleanup that produced this layout, see `docs/refactor_summary.md`.
