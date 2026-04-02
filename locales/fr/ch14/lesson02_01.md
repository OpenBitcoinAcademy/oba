## Depenser avec la cle modifiee

La depense par chemin de cle est la maniere la plus elegante et la plus privee de depenser une sortie Taproot. Le temoin contient un seul element : une signature Schnorr BIP 340 de 64 octets pour la cle modifiee Q.

Le depenseur calcule la cle privee modifiee (cle privee originale + tweak), signe la transaction et fournit la signature. Aucun script n'est revele. Aucune preuve de Merkle n'est necessaire. Le temoin entier fait 64 octets.

Un verificateur verifie la signature par rapport a Q (la cle dans le scriptPubKey de la sortie). Si elle est valide, la depense est autorisee. Le verificateur ne sait pas et ne se soucie pas de savoir si Q a ete derive d'une cle unique, d'un agregat MuSig2 ou d'une cle avec un arbre de scripts embarque.
