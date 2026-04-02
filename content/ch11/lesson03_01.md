## What is a Merkle Tree?

A Merkle tree is a binary tree of hashes. Each leaf node is the hash of a transaction. Each internal node is the hash of its two children concatenated. The root of the tree (the Merkle root) is a single hash that commits to every transaction in the block.

If a block contains four transactions (A, B, C, D), the tree is: hash(A) and hash(B) combine into hash(AB). hash(C) and hash(D) combine into hash(CD). hash(AB) and hash(CD) combine into the Merkle root.

If the number of transactions is odd, the last hash is duplicated to make an even count at each level.
