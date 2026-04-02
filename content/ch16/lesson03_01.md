## The Three Maps

A PSBT contains three types of key-value maps: one global map, one map per input, and one map per output. Each map stores typed key-value pairs terminated by a 0x00 byte. The key type determines what the entry means. The key payload and value payload carry the data.

The global map holds data that applies to the entire transaction. In PSBTv0, the most important global entry is UNSIGNED_TX (key type 0x00), which contains the full unsigned transaction. In PSBTv2, the global map instead holds TX_VERSION, FALLBACK_LOCKTIME, INPUT_COUNT, and OUTPUT_COUNT as separate fields. The unsigned transaction is reconstructed from per-input and per-output maps rather than stored whole.

Global entries also include XPUB (key type 0x01), which maps an extended public key to its BIP 32 derivation path. This lets a signer verify that change outputs derive from the same wallet root without needing access to the full wallet descriptor.
