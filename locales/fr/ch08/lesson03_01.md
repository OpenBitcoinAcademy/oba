## Schnorr : une signature plus elegante

Les signatures Schnorr (BIP340) ont ete activees avec la mise a niveau Taproot en novembre 2021. Elles sont utilisees pour les transactions segwit v1 (P2TR).

L'algorithme Schnorr est mathematiquement plus elegant qu'ECDSA. Il produit une signature fixe de 64 octets (contre 71-73 octets variables pour ECDSA). L'equation de verification est lineaire, ce qui permet l'agregation de signatures : plusieurs signatures peuvent etre combinees en une seule signature de la meme taille.

Les signatures Schnorr ont une reduction de securite prouvable au probleme du logarithme discret, une garantie de securite plus forte qu'ECDSA.
