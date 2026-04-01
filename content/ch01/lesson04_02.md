## The Derivation Steps

Creating a Bitcoin address from a public key requires several hashing steps.

First, hash the public key with SHA-256. Then hash that result with RIPEMD-160. This produces a 20-byte hash called the public key hash, or HASH160.

Next, add a version byte to the front. This byte identifies the network (mainnet vs testnet) and the address type.

Finally, encode the result. Legacy addresses use Base58Check encoding, which adds a 4-byte checksum and encodes in a human-readable alphabet that avoids confusing characters like 0/O and l/1. Newer SegWit addresses use Bech32 encoding, which is lowercase-only and has stronger error detection.
