## Why Tweaking Works

The tweak binds the script tree to the public key cryptographically. Changing any script in the tree changes the Merkle root, which changes the tweak, which changes Q. A script committed in Q at output creation time cannot be modified later.

The private key is tweaked identically: tweaked_privkey = (privkey + t) mod n. The holder of the internal private key can compute the tweaked private key and sign directly. This is the key path spend.

Tagged hashes prevent collisions between different uses of the hash function. The tag "TapTweak" is domain-separated: SHA256(SHA256("TapTweak") || SHA256("TapTweak") || data). Different tags produce different hash outputs for the same data, eliminating cross-protocol attacks.
