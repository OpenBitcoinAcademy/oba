## What is a Block?

A block is a batch of transactions bundled together and added to the blockchain. Each block contains a header and a list of transactions.

The block header has six fields: version, previous block hash, merkle root (a hash of all transactions in the block), timestamp, difficulty target, and nonce. The previous block hash links each block to the one before it, forming a chain.

Miners compete to find a nonce value that makes the block header hash fall below the difficulty target. This requires trillions of guesses. The first miner to find a valid nonce broadcasts the block. Other nodes verify it and add it to their copy of the chain.
