## MuSig2 as Internal Key

The internal key P can be a MuSig2 aggregate of multiple public keys. If Alice and Bob aggregate their keys into P using MuSig2, the tweaked key Q commits to both of them. When they cooperate, they produce a single Schnorr signature on Q.

On-chain, this 2-of-2 multisig looks identical to a single-signer transaction. The output is 34 bytes. The witness is 64 bytes. No observer can determine that two parties were involved.

BitGo reported that a MuSig2 key path input costs 57.5 virtual bytes, compared to 104.5 vbytes for a native SegWit P2WSH multisig input. The savings come from eliminating the per-signer public keys and signatures that multisig scripts require on-chain.
