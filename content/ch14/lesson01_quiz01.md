How is the tweaked public key Q computed?

- Q = SHA256(P)
- Q = P * merkle_root
- Q = P + tagged_hash("TapTweak", P || merkle_root) * G
