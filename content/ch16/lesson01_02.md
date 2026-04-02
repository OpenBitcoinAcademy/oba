## What a PSBT Is

A Partially Signed Bitcoin Transaction (PSBT) is a binary format defined by BIP 174. It wraps an unsigned transaction together with the metadata each participant needs to do its job. Creators attach UTXO data. Updaters add derivation paths. Signers add signatures. The format carries everything needed so that each role can act independently.

The binary begins with a five-byte magic: `0x70736274FF`. The first four bytes spell "psbt" in ASCII. The fifth byte, `0xFF`, marks the separator. Any tool receiving a blob can check these five bytes to confirm it is a PSBT before parsing further.

After the magic comes a sequence of key-value maps. Each map entry has a key type, a key payload, and a value payload. A zero byte (0x00) terminates each map. The first map is the global map. Subsequent maps alternate between per-input and per-output maps, one for each input and output in the transaction.
