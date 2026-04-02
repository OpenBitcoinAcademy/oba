## Mining as a Decentralized Lottery

Bitcoin has no central authority that decides which transactions go into blocks. Instead, miners compete in a computational lottery. The winner earns the right to add the next block to the chain.

Each miner takes a set of unconfirmed transactions, assembles them into a candidate block, and hashes the block header using SHA-256 (applied twice). The result must be a number below the current target. Most attempts fail. The miner changes the nonce field in the header and hashes again.

This process repeats billions of times per second across all miners worldwide. The miner who finds a valid hash first broadcasts the block. Other nodes verify it independently. No one needs permission to mine, and no one can predict who will win next.
