## The Wire Format

A Bitcoin transaction is serialized into a sequence of bytes for transmission across the network and storage in blocks. The serialization format defines the exact byte order and encoding of every field.

A legacy transaction has four top-level fields serialized in order: version (4 bytes, little-endian), inputs (variable), outputs (variable), and locktime (4 bytes, little-endian).

A segwit transaction adds three fields: after the version, a marker byte (0x00) and a flag byte (0x01) signal that witness data follows. The witness structure appears after the outputs and before the locktime. A parser that sees a zero byte where the input count should be knows it is reading the extended (segwit) format.
