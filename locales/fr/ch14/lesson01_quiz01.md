Comment la cle publique modifiee Q est-elle calculee ?

- Q = SHA256(P)
- Q = P * merkle_root
- Q = P + tagged_hash("TapTweak", P || merkle_root) * G
