## Per-Input Fields

Each input map describes one input of the transaction. The critical fields a signer needs are the UTXO data, the signing key derivation path, and any scripts required to construct the witness.

WITNESS_UTXO (key type 0x01) stores the output being spent: the amount in satoshis and the scriptPubKey. For SegWit inputs, this is sufficient for signing because the sighash algorithm commits to the amount directly. NON_WITNESS_UTXO (key type 0x00) stores the entire previous transaction, necessary for legacy inputs where the signer must verify the amount by looking up the output in the full transaction.

BIP32_DERIVATION (key type 0x06) maps a public key to its BIP 32 derivation path and master key fingerprint. The signer matches the fingerprint against its own master key, derives the private key at the given path, and signs. PARTIAL_SIG (key type 0x02) stores a signature keyed by the public key that produced it. SIGHASH_TYPE (key type 0x03) specifies which sighash flag the signer should use.

For P2SH inputs, REDEEM_SCRIPT (key type 0x04) carries the redeem script. For P2WSH inputs, WITNESS_SCRIPT (key type 0x05) carries the witness script. The signer needs these to compute the correct sighash.
