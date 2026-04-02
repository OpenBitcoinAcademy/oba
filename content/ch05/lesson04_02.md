## The Optional Passphrase

BIP39 supports an optional passphrase that is mixed into the seed derivation. The same 12 words with different passphrases produce different seeds and different wallets.

This provides plausible deniability: a user can create a decoy wallet with one passphrase and a real wallet with another. Under coercion, they can reveal the decoy passphrase without exposing the main funds.

The risk: forgetting the passphrase means losing access to the wallet. There is no recovery mechanism. The passphrase is not stored anywhere. It exists only in the user's memory or physical backup.
