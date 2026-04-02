## What Output Descriptors Are

An output descriptor is a string that fully describes how to derive a set of Bitcoin addresses. BIPs 380 through 389 define the descriptor language. Each descriptor specifies the script type, the keys involved, and the derivation paths.

A descriptor like `wpkh([d34db33f/84h/0h/0h]xpub.../0/*)` tells a wallet everything it needs: use P2WPKH, derive from this extended public key at this path, and generate sequential addresses. The wallet does not need to guess the script type or derivation scheme. The descriptor is self-contained.

Before descriptors, wallets relied on conventions. BIP 44 said HD wallets should use a specific derivation path. BIP 49 added a different path for P2SH-SegWit. BIP 84 added another for native SegWit. A wallet importing an xpub had to try all conventions and hope it guessed correctly. Descriptors replace guesswork with explicit declarations.
