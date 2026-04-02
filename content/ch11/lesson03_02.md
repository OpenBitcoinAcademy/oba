## Merkle Proofs

A Merkle proof allows verification that a specific transaction is included in a block without downloading every transaction. The proof consists of the transaction hash, the sibling hashes along the path to the root, and the Merkle root from the block header.

To verify, start with the transaction hash. Hash it with its sibling. Hash the result with the next sibling. Continue until you reach the root. If your computed root matches the Merkle root in the block header, the transaction is in the block.

A block with 4,000 transactions requires a Merkle proof of only about 12 hashes (log2(4000)). SPV clients use Merkle proofs to verify transactions without downloading full blocks.
