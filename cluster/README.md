# Cluster

`cluster/` contains the distributed data plane layered on top of the storage engine.

- `config/`: shared topology schema
- `percolator/`: MVCC and two-phase commit logic
- `raft/`: raft implementation and tests
- `raftstore/`: region lifecycle, transport, client, and server integration
