# Refactor Summary

This document records the repository-structure refactor implemented on branch `refactor-clean-root-layout`.

It is not an architecture document. It is a change log for the layout cleanup so new contributors can see what moved and why.

## Goals

- remove the flat, mixed-purpose repository root
- group code by runtime domain instead of by historical package accretion
- separate shipping binaries, implementation code, tooling, and black-box tests
- keep import paths and docs consistent after the move
- preserve buildability and test coverage during the refactor

## High-level result

Before the refactor, the repository mixed engine files, tests, storage packages, distributed packages, and helper scripts at the top level.

After the refactor:

- the repository root contains no Go source files
- the embedded API lives under `engine/`
- local storage code lives under `storage/`
- distributed data-plane code lives under `cluster/`
- PD-lite lives under `pd/`
- observability code lives under `observability/`
- scripts and benchmarks live under `tools/`
- black-box and layout tests live under `tests/`

## Major moves

| Old location | New location | Notes |
| --- | --- | --- |
| root `db.go`, `options.go`, `iterator.go`, `stats.go`, `vlog.go`, write pipeline files | `engine/` | Embedded public API and engine wiring were moved out of the root. |
| `kv/` | `storage/kv/` | Entry format, internal keys, column families. |
| `lsm/` | `storage/lsm/` | Memtables, SSTables, flush, compaction, iterators. |
| `manifest/` | `storage/manifest/` | Durable metadata and version edits. |
| `wal/` | `storage/wal/` | Typed WAL and replay. |
| `vlog/` | `storage/vlog/` | Value-log manager and IO. |
| `file/` | `storage/file/` | mmap-backed file helpers. |
| `vfs/` | `storage/vfs/` | Filesystem abstraction and fault injection. |
| `utils/cache/` | `storage/cache/` | Cache primitives promoted into storage domain. |
| `utils/mmap/` | `storage/mmap/` | mmap platform helpers promoted into storage domain. |
| `config/` | `cluster/config/` | Shared cluster topology schema. |
| `percolator/` | `cluster/percolator/` | MVCC and 2PC logic. |
| `raft/` | `cluster/raft/` | Raft implementation. |
| `raftstore/` | `cluster/raftstore/` | Region lifecycle, transport, client, server. |
| `pd/` after temporary nesting under `controlplane/pd/` | `pd/` | Final flattened control-plane location. |
| `metrics/` | `observability/metrics/` | Metrics counters and snapshots. |
| `hotring/` | `observability/hotring/` | Hot-key tracking. |
| `benchmark/` | `tools/bench/` | Separate benchmark module kept under tools. |
| `scripts/` | `tools/scripts/` | Local cluster, recovery, proto, and benchmark scripts. |
| temporary `internal/utils/` | `utils/` | Final flattened shared helper package. |

## Root cleanup

The root package `github.com/feichai0017/NoKV` was kept briefly as a compatibility shim during the move to `engine/`.

That shim was then removed so the repository root now acts only as:

- module metadata
- top-level docs
- CI and build config
- entrypoint directories such as `cmd/`, `engine/`, `storage/`, and `cluster/`

The layout guard in `tests/layout/repo_layout_test.go` enforces that no top-level `.go` files are reintroduced.

## Test organization changes

The refactor did not force all tests into one directory. The final rule is:

- black-box and repository-shape tests go under `tests/`
- white-box tests stay beside implementation when they need package-private access

Current examples:

- `tests/facade/` validates the embedded API surface
- `tests/layout/` enforces repository shape
- package-internal algorithm and subsystem tests remain beside `storage/*`, `cluster/*`, `pd/*`, `utils/`, and `engine/`

## Tooling and docs updates

These repo-level surfaces were updated to match the new layout:

- `README.md`
- `docs/*`
- `Makefile`
- `Dockerfile`
- `tools/scripts/*`
- `tools/bench/go.mod`

Notable command-path changes:

- `./scripts/run_local_cluster.sh` -> `./tools/scripts/run_local_cluster.sh`
- `./scripts/run_benchmarks.sh` -> `./tools/scripts/run_benchmarks.sh`
- benchmark module path -> `github.com/feichai0017/NoKV/tools/bench`
- embedded import path -> `github.com/feichai0017/NoKV/engine`

## Intentional non-moves

Some top-level directories were intentionally kept where they are:

- `cmd/` stays top-level because it is the standard location for binaries
- `pb/` stays top-level because generated protobuf packages are shared broadly
- `docs/` stays top-level because it describes the whole repository

## Verification run on this branch

The following verification was run successfully after the refactor:

```bash
go test ./tests/layout ./tests/facade ./engine ./storage/... ./cluster/config ./cluster/percolator/... ./pd/... ./observability/... ./cluster/raft ./utils ./cmd/nokv-config ./cluster/raftstore/command ./cluster/raftstore/engine ./cluster/raftstore/failpoints

go test -run '^$' ./cmd/... ./cluster/raftstore/...

go -C tools/bench test -run '^$' ./...
```

Notes:

- full runtime tests for listener-heavy packages were compile-only in this environment because the sandbox restricts local TCP binding
- the benchmark module was compile-checked, not executed as a workload run

## Current target layout

For the current canonical directory map, see `docs/project_layout.md`.
