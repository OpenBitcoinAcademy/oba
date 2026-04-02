## The Chain of Hashes

Each block header contains the hash of the previous block's header. Block 500,000 contains the hash of block 499,999, which contains the hash of block 499,998, and so on back to the genesis block (block 0).

Changing any data in block 499,999 would change its hash. Block 500,000 would then contain an incorrect previous block hash and become invalid. To alter a historical block, an attacker must redo the proof of work for that block and every block after it.

This cumulative linking is what makes the blockchain tamper-evident. The deeper a block is, the more work is required to alter it. Each new block adds another layer of protection to every block beneath it.
