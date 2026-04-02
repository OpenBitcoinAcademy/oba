## Byte by Byte

**Version** (4 bytes): 01000000 means version 1. 02000000 means version 2 (enables BIP68 sequence constraints). Little-endian encoding: the least significant byte comes first.

**Inputs**: a varint count followed by each input. Each input contains the previous transaction hash (32 bytes, reversed), the output index (4 bytes), the input script length (varint), the input script (variable), and the sequence number (4 bytes).

**Outputs**: a varint count followed by each output. Each output contains the value in satoshis (8 bytes, little-endian) and the output script length (varint) followed by the output script.

**Witness** (segwit only): for each input, a count of witness items followed by each item's length and data. Legacy inputs have zero witness items.

**Locktime** (4 bytes): usually 00000000. A nonzero value restricts when the transaction can be mined.

The transaction ID (txid) is the double-SHA-256 hash of the serialized transaction, with the witness data excluded (for segwit) or included (for legacy).
