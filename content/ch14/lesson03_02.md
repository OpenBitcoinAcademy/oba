## Building the Script Tree

The Merkle tree is built from TapLeaf and TapBranch hashes.

Each leaf: tagged_hash("TapLeaf", leaf_version || compact_size(script_length) || script). The leaf version for current Tapscript is 0xC0.

Each branch: tagged_hash("TapBranch", sorted(left_hash, right_hash)). The two children are sorted lexicographically before hashing, ensuring the tree has a single canonical form regardless of insertion order.

The tree can have up to 128 levels of depth, allowing up to 2^128 leaves. Practical trees are far smaller. Scripts placed at shallow depths (closer to the root) have shorter Merkle proofs and cost less to spend. Place the most likely fallback at depth 1. Place rarely-used emergency paths deeper.
