## Identifiable Aborts and MuSig2 vs FROST

A malicious participant can disrupt signing by submitting an invalid signature share. Without additional checks, the coordinator would sum the shares and produce an invalid final signature, but would not know which signer misbehaved. FROST supports identifiable aborts: the coordinator can verify each signature share individually against the signer's public key share. The misbehaving signer is identified and can be excluded from future sessions.

MuSig2 and FROST serve different needs. MuSig2 is an n-of-n scheme: all participants must sign, there is no threshold, and the protocol is simpler. FROST is a t-of-n scheme: it tolerates absent signers but requires a more complex key generation phase. Both produce identical on-chain output: a single 64-byte Schnorr signature under one 32-byte public key.

MuSig2 suits scenarios where all signers are expected to be available, such as a Lightning channel between two nodes. FROST suits custody arrangements where redundancy matters, such as a corporate treasury with five keyholders where any three can authorize a spend.
