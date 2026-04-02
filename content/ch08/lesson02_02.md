## DER Serialization

ECDSA signatures in Bitcoin are encoded using DER (Distinguished Encoding Rules), a standard binary format. The DER encoding wraps the r and s values with type and length bytes.

A typical DER-encoded ECDSA signature is 71 to 73 bytes long. It is followed by a one-byte SIGHASH flag that indicates which parts of the transaction the signature commits to.

Segwit v0 transactions still use ECDSA but require strict DER encoding (BIP66). This eliminated a source of transaction malleability where the same valid signature could be encoded in multiple ways, producing different transaction IDs.
