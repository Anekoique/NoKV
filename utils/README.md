# Utils

`utils/` contains shared low-level helpers still used across the engine, storage, and cluster layers.

It includes structures such as ART/skiplist indexes, iterators, pools, throttles, random helpers, and common sentinels.

The package is intentionally flat for now because these helpers are still consumed widely across the repository.
