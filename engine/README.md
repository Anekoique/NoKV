# Engine

This package holds the embeddable NoKV database API and runtime.

Why it exists:

- gives the embedded database a clear, explicit import path
- keeps the module root free of implementation details
- keeps engine-specific unit tests next to the implementation they exercise

This is the intended package for embedders. The module-root `NoKV` package is
now only a temporary compatibility shim during the broader repo refactor.
