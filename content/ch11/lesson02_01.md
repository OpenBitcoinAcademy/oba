## Block Header Structure

Every block begins with an 80-byte header containing six fields.

**Version** (4 bytes): indicates which consensus rules the block follows and signals support for soft fork proposals via versionbits.

**Previous Block Hash** (32 bytes): the double-SHA-256 hash of the previous block's header. This links every block to its predecessor, forming the chain.

**Merkle Root** (32 bytes): a single hash that commits to every transaction in the block. Changing any transaction changes the merkle root.

**Timestamp** (4 bytes): the approximate time the block was mined, in Unix epoch seconds.

**Target Bits** (4 bytes): a compact encoding of the proof-of-work target. The block header hash must be below this target.

**Nonce** (4 bytes): the value miners change to search for a valid hash.
