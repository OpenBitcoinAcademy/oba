## Resource Requirements

Running a full node requires disk space, bandwidth, and time. The initial blockchain download is over 500 GB as of 2024 and grows by roughly 150 MB per day.

Bitcoin Core can run in pruned mode, discarding old block data after verification and keeping only the UTXO set and recent blocks. This reduces storage to under 10 GB. The node still verifies the entire history during initial sync.

A full node can run on hardware as modest as a Raspberry Pi. The initial sync takes longer on slow hardware, but once synchronized, the daily data is manageable on most home internet connections.
