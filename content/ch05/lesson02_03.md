## From Seeds to Trees

A seed is a random number, typically 128 to 256 bits. It is the root of all key derivation. From this seed, a hierarchy of keys can be generated: the seed produces a master key, the master key produces child keys, each child can produce its own children.

This tree structure is the basis of Hierarchical Deterministic (HD) wallets, defined in BIP32. The tree allows organizing keys by purpose: one branch for receiving payments, another for change, another for a different account.

The seed never changes. Every key in the tree can be regenerated from it. A single backup protects an unlimited number of future keys.
