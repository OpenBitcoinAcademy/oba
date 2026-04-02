## BIP32: The HD Wallet Standard

BIP32 defines how to derive a tree of keys from a single seed. The process uses HMAC-SHA512 to split each derivation step into a child key and a chain code. The chain code is mixed into the next derivation, preventing anyone from computing sibling keys.

The master key sits at the root of the tree. From it, a sequence of child keys is derived, each identified by an index number (0, 1, 2, ...). Each child can produce its own children, forming a tree of arbitrary depth.

This tree structure gives wallets organizational power. Different branches of the tree serve different purposes, and the entire tree is recoverable from the original seed.
