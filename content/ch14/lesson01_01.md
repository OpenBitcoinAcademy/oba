## The Internal Key and the Tweak

A Taproot output locks funds to a tweaked public key Q. This key is derived from two inputs: an internal public key P and an optional Merkle root of a script tree.

The tweak value t is computed as: t = tagged_hash("TapTweak", P || merkle_root). The tweaked key is: Q = P + t * G, where G is the generator point on secp256k1.

If there is no script tree, the tweak uses only P: t = tagged_hash("TapTweak", P). The output still commits to the internal key, but no scripts are embedded.

On-chain, the scriptPubKey is: OP_1 followed by the 32-byte X coordinate of Q. This is 34 bytes, the same size as a P2WSH output. No observer can tell whether Q contains an embedded script tree or not.
