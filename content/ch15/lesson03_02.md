## Descriptors and Miniscript Together

Descriptors support embedding Miniscript expressions. The `wsh()` descriptor wraps a Miniscript expression in a P2WSH output. The `tr()` descriptor places a Miniscript expression inside a Taproot script tree leaf.

A descriptor like `wsh(and_v(v:pk(Alice),or_d(pk(Bob),older(12960))))` defines a complete spending policy as a single string. Any wallet that understands the descriptor format can import this string, derive the correct addresses, identify the spending conditions, and construct valid transactions. The descriptor carries the full policy, not a partial hint.

This is what makes wallets interoperable. A hardware signer can display the spending conditions parsed from the descriptor. A watch-only wallet can monitor the resulting addresses. A signing coordinator can construct a PSBT that satisfies the policy. Each tool reads the same descriptor string and arrives at the same Script output.

Descriptors include a checksum: eight characters appended after a `#` symbol. The checksum detects transcription errors. A wallet rejects any descriptor whose checksum does not match. A single changed character in the key or policy produces a different checksum, catching mistakes before funds are sent to the wrong address.
