## The Multisig Landscape

Bitcoin has supported multi-party custody since 2012 through OP_CHECKMULTISIG. A 2-of-3 multisig script lists all three public keys on-chain and requires two valid signatures. This works, but it reveals exactly how many parties control the funds. Anyone inspecting the blockchain sees the threshold, the total number of signers, and every public key involved.

The on-chain cost scales linearly. A 3-of-5 multisig places five 33-byte public keys and three 72-byte signatures into the witness. That is 381 bytes of witness data before counting the script itself. A 7-of-10 setup costs even more. Fees rise, privacy drops, and the spending policy is visible to all observers.

Taproot and Schnorr signatures changed what is possible. MuSig2 allows n participants to produce a single aggregated signature that looks identical to a solo Schnorr signature on-chain. The combined public key is 32 bytes. The signature is 64 bytes. No observer can tell how many parties were involved.
