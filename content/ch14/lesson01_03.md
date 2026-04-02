## X-Only Public Keys

Taproot uses 32-byte X-only public keys instead of the traditional 33-byte compressed format. The Y coordinate is implicitly chosen to be even. This saves one byte per key on-chain and in signatures.

A standard compressed public key has a prefix byte (0x02 for even Y, 0x03 for odd Y) followed by 32 bytes of X coordinate. Taproot drops the prefix and requires the Y coordinate to be even. If the computed Q has an odd Y, the implementation negates the private key (which flips Y to even) before signing.

This convention simplifies batch verification and key aggregation. Every Taproot output key is 32 bytes. Every Schnorr signature is 64 bytes. There are no variable-length encodings.
