## Taproot Fields and Version Differences

BIP 371 added Taproot-specific fields to the PSBT format. TAP_KEY_SIG (input key type 0x13) stores a Schnorr signature for a key path spend. TAP_SCRIPT_SIG (input key type 0x14) stores a Schnorr signature for a specific script leaf, keyed by both the X-only public key and the leaf hash.

TAP_LEAF_SCRIPT (input key type 0x15) provides the script, its leaf version, and the control block needed for script path spending. TAP_BIP32_DERIVATION (input key type 0x16) extends the standard BIP 32 derivation field with a list of leaf hashes the key can sign for. TAP_INTERNAL_KEY (input key type 0x17) records the internal public key before tweaking.

On the output side, TAP_INTERNAL_KEY (output key type 0x05) and TAP_BIP32_DERIVATION (output key type 0x07) let signers verify that Taproot change outputs belong to the same wallet. The signer can re-derive the tweaked key from the internal key and confirm it matches the output's scriptPubKey.

PSBTv2 (BIP 370) differs from v0 in structure, not in concept. In v0, the unsigned transaction lives in the global map as a single serialized blob. In v2, inputs and outputs are described by per-map fields: PREVIOUS_TXID, OUTPUT_INDEX, SEQUENCE for inputs; AMOUNT and SCRIPT for outputs. Constructors can add inputs and outputs incrementally without re-serializing the entire transaction each time.
