## Privacy Through Selective Disclosure

A Taproot output with 8 script leaves reveals only one leaf when spent via script path. The other 7 leaves appear as hashes in the Merkle proof. An observer learns one spending condition but cannot determine what the other conditions were.

Compare this to P2WSH multisig: spending reveals the full script, all public keys, and which keys signed. Every party involved is visible.

A 2-of-3 multisig using Taproot: the key path holds the MuSig2 aggregate of the two most common signers. Two script leaves hold the fallback pairs (A+C and B+C). In the common case, the key path is used and nothing about the multisig is revealed. In a fallback, only one alternative pair is exposed. The other pair remains a hash.
