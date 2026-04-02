## Wrapped SegWit (P2SH-P2WPKH)

Not all wallets supported native segwit addresses (bc1...) when segwit was first activated. Wrapped segwit provides backward compatibility by placing the segwit program inside a P2SH redeem script.

The output looks like a standard P2SH address (starting with "3"). The redeem script contains: OP_0 <20-byte-pubkey-hash>. When spending, the input script reveals the redeem script, and the witness provides the signature and public key.

Wrapped segwit is a transitional format. Native segwit (bc1q... addresses) is preferred for new transactions because it is smaller, cheaper, and has better error detection via bech32 encoding.
