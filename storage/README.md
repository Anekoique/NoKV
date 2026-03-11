# Storage

`storage/` contains the local persistence and indexing layers used by both embedded and clustered NoKV.

- `kv/`: entry formats, key encoding, column families
- `lsm/`: memtables, SSTables, flush, compaction, iterators
- `manifest/`: durable metadata and version edits
- `wal/`: typed write-ahead log and replay
- `vlog/`: value-log segments, IO, and GC support
- `file/`: mmap-backed file helpers
- `vfs/`: filesystem abstraction and fault injection
- `cache/`: cache primitives
- `mmap/`: mmap platform helpers
