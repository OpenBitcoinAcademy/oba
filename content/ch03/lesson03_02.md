## Exploring the Blockchain

With `bitcoin-cli`, you can retrieve any block by its height or hash. The `getblock` command returns the block header fields, the list of transaction IDs, and metadata like the number of confirmations.

The `getrawtransaction` command returns a transaction's raw hex data or its decoded JSON representation, showing inputs, outputs, scripts, and witness data.

The `getblockchaininfo` command reports the current chain height, difficulty, verification progress, and whether the node is still synchronizing. These commands turn abstract blockchain concepts into concrete, inspectable data.
