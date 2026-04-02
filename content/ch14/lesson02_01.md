## Spending with the Tweaked Key

Key path spending is the simplest and most private way to spend a Taproot output. The witness contains a single element: a 64-byte BIP 340 Schnorr signature for the tweaked key Q.

The spender computes the tweaked private key (original private key + tweak), signs the transaction, and provides the signature. No scripts are revealed. No Merkle proofs are needed. The entire witness is 64 bytes.

A verifier checks the signature against Q (the key in the output's scriptPubKey). If valid, the spend is authorized. The verifier does not know or care whether Q was derived from a single key, a MuSig2 aggregate, or a key with an embedded script tree.
