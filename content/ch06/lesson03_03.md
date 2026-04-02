## Varints and Endianness

Bitcoin uses two encoding conventions throughout its serialization.

**Varints** (variable-length integers) encode counts and lengths. A value below 253 fits in one byte. Values 253-65535 use a 0xFD prefix followed by 2 bytes. Values up to ~4 billion use 0xFE followed by 4 bytes. Larger values use 0xFF followed by 8 bytes.

**Little-endian** byte order places the least significant byte first. The version 1 is serialized as 01 00 00 00, not 00 00 00 01. Satoshi values, block heights, and most numeric fields use little-endian.

Transaction hashes (txids) are displayed in reversed byte order by convention. The double-SHA-256 hash produces bytes in one order, but block explorers and Bitcoin Core display them reversed. This is a display convention, not a serialization rule.
