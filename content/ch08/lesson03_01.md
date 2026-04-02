## Schnorr: A Simpler Signature

Schnorr signatures (BIP340) were activated with the Taproot upgrade in November 2021. They are used for segwit v1 (P2TR) transactions.

The Schnorr algorithm is mathematically simpler than ECDSA. It produces a fixed 64-byte signature (compared to ECDSA's variable 71-73 bytes). The verification equation is linear, which enables signature aggregation: multiple signatures can be combined into a single signature of the same size.

Schnorr signatures have a provable security reduction to the discrete logarithm problem, a stronger security guarantee than ECDSA.
