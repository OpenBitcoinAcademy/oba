## Target, Difficulty, and the Nonce

The target is a 256-bit number. A valid block hash must be numerically less than the target. A lower target means fewer valid hashes exist, which makes finding one harder. Difficulty is the inverse of the target: when the target goes down, difficulty goes up.

The block header contains a 32-bit nonce field. Miners increment this value with each hash attempt, cycling through all $2^{32}$ (about 4.3 billion) possibilities. Modern miners exhaust the nonce space in under a second.

When the nonce space runs out, miners modify other fields to create new hash inputs. The most common technique changes the coinbase transaction's extra nonce field. This changes the merkle root in the block header, giving the miner a fresh set of $2^{32}$ nonces to try.
