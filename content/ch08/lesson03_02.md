## Multisignature Aggregation

With ECDSA, a transaction requiring three signers needs three separate signatures, each 71-73 bytes. The signatures are verified individually.

With Schnorr, the same three signers can combine their signatures into a single 64-byte aggregate signature using protocols like MuSig2. On the blockchain, the aggregate looks identical to a single-signer signature. No observer can tell how many parties participated.

This has two benefits. Transactions with multiple signers use less block space (lower fees). And multisignature transactions become indistinguishable from single-signature transactions, improving privacy.
