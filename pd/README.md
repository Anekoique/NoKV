# PD

`pd/` contains PD-lite, the runtime control plane for clustered NoKV.

- `client/`: PD client helpers
- `core/`: cluster state and ID allocation
- `server/`: gRPC services and schedulers
- `storage/`: local persistence
- `tso/`: timestamp allocation
- `adapter/`: integration helpers
