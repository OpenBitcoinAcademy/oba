## From Shares to a Standard Signature

The coordinator collects t signature shares and sums them. The result is a single Schnorr signature: a 64-byte (R, s) pair, where R is the aggregate nonce point and s is the sum of signature shares. This signature is valid under the group's aggregate public key.

A Bitcoin full node verifying this transaction runs the standard BIP 340 Schnorr verification equation. It checks one signature against one public key. The node has no way to know that the signature was produced by three signers, five signers, or one signer. The verification is identical.

This is why FROST is powerful for Bitcoin: it requires no consensus changes. Taproot and BIP 340 already accept the signatures FROST produces. The threshold signing complexity lives entirely in the signers' software. The blockchain and all verifiers remain unaware.
