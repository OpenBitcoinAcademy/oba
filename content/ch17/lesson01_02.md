## From n-of-n to t-of-n

MuSig2 has one limitation: it requires all n participants to sign. Every keyholder must be online and cooperating at signing time. If one participant loses their key or goes offline, the funds are stuck. This makes MuSig2 a pure n-of-n scheme.

FROST (Flexible Round-Optimized Schnorr Threshold signatures) solves this. FROST is a threshold signature scheme: any t of n participants can produce a valid signature. A 3-of-5 FROST setup means any three of the five keyholders can sign. The other two can be offline, unavailable, or even permanently lost.

The signature that FROST produces is a standard 64-byte Schnorr signature. On-chain, it is indistinguishable from a single-key spend. The threshold policy, the number of signers, and the individual public keys are all hidden. A FROST spend looks exactly like a key path spend from a single Taproot address.
