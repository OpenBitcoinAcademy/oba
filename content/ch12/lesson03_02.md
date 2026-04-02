## Block Validation Rules

When a node receives a new block, it checks dozens of criteria before accepting it. The block header hash must be below the current target. The timestamp must be within acceptable bounds. The first transaction must be a coinbase, and the coinbase output must not exceed the allowed subsidy plus fees. Every other transaction must be valid according to the script rules.

The block's merkle root must match the hash tree of all included transactions. The block size must not exceed the consensus limit. The block must reference a valid parent block that the node already has.

If any check fails, the node rejects the block and may disconnect from the peer that sent it. There is no appeal process. A block is either valid or it is not. This strict validation is what prevents miners from creating bitcoin out of thin air or spending coins they do not own.
