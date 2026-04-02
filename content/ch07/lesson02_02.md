## Multisignature Scripts

A multisignature (multisig) script requires M signatures from N public keys to authorize spending. A 2-of-3 multisig requires any 2 of 3 designated keys.

In a bare multisig output, the locking script contains: M <pubkey1> <pubkey2> ... <pubkeyN> N OP_CHECKMULTISIG. The unlocking script provides: OP_0 <sig1> <sig2>. The leading OP_0 is a workaround for a historical off-by-one bug in OP_CHECKMULTISIG.

In practice, multisig is wrapped in P2SH. The sender sees a standard P2SH address. The multisig details are hidden inside the redeem script and revealed only when spending. This keeps outputs compact and the complexity private until spending time.
