## Signatures in Bitcoin Transactions

In Bitcoin, signatures appear in transaction inputs (for legacy transactions) or in the witness structure (for segwit transactions). Each input that spends a UTXO requires a signature proving the spender controls the key that locked that output.

Multiple parties can collaborate on a single transaction. Each party signs only the input they control. The signatures are independent: one signer does not need access to another's private key.

Bitcoin uses two signature algorithms: ECDSA (Elliptic Curve Digital Signature Algorithm) for legacy and segwit v0 transactions, and Schnorr signatures for segwit v1 (Taproot) transactions.
