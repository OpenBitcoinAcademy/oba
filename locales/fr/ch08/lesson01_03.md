## Signatures dans les transactions Bitcoin

Dans Bitcoin, les signatures apparaissent dans les entrees de transaction (pour les transactions legacy) ou dans la structure temoin (pour les transactions segwit). Chaque entree qui depense un UTXO necessite une signature prouvant que le depenseur controle la cle qui a verrouille cette sortie.

Plusieurs parties peuvent collaborer sur une seule transaction. Chaque partie signe uniquement l'entree qu'elle controle. Les signatures sont independantes : un signataire n'a pas besoin d'acceder a la cle privee d'un autre.

Bitcoin utilise deux algorithmes de signature : ECDSA (Elliptic Curve Digital Signature Algorithm) pour les transactions legacy et segwit v0, et les signatures Schnorr pour les transactions segwit v1 (Taproot).
