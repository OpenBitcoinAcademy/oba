## The Data Directory

Bitcoin Core stores its data in a platform-specific directory: `~/.bitcoin` on Linux, `~/Library/Application Support/Bitcoin` on macOS, `%APPDATA%\Bitcoin` on Windows.

The `blocks/` subdirectory contains the raw block data files. The `chainstate/` directory holds the UTXO database, a LevelDB store of all unspent outputs. The `mempool.dat` file persists the mempool across restarts.

The configuration file `bitcoin.conf` controls network settings, resource limits, RPC authentication, and feature flags. Settings like `prune=550` enable pruned mode, and `txindex=1` builds an index of all transactions for faster lookups.
