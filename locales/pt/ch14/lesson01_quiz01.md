Como a chave publica ajustada Q e calculada?

- Q = SHA256(P)
- Q = P * merkle_root
- Q = P + tagged_hash("TapTweak", P || merkle_root) * G
