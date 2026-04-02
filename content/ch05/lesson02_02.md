## Deterministic Key Generation

Deterministic key generation solves the backup problem. A single random seed produces a sequence of private keys through repeated hashing. Every key in the sequence can be recreated from the seed.

Back up the seed once, and you can recover every key the wallet has generated or will generate. The seed is the master secret. Lose it and you lose access to all derived keys.

The simplest form: start with a seed, hash it to get key 1, hash key 1 to get key 2, and so on. This produces a flat list of keys. But modern wallets use a more sophisticated approach: a tree.
