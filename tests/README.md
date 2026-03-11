# Tests Layout

- `tests/facade/` contains black-box tests for the public
  `github.com/feichai0017/NoKV/engine` package surface.
- package-local tests remain next to implementation code when they need access
  to unexported helpers or internal state.

This split keeps the module root clean while preserving idiomatic Go testing for
internal behavior.
