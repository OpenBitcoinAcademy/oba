## Chemins de derivation (BIP44)

BIP44 definit une structure standard pour l'arbre de cles. Le format du chemin est : m / purpose' / coin_type' / account' / change / address_index.

**purpose** est 44 pour les adresses standard (ou 84 pour SegWit, 86 pour Taproot). **coin_type** est 0 pour le mainnet Bitcoin, 1 pour le testnet. **account** permet de separer les fonds en groupes logiques. **change** est 0 pour les adresses de reception et 1 pour les adresses de monnaie. **address_index** s'incremente pour chaque nouvelle adresse.

Un chemin typique d'adresse de reception Bitcoin : m/84'/0'/0'/0/0. Cela signifie : objectif SegWit, mainnet Bitcoin, premier compte, reception (pas monnaie), premiere adresse.

Les chemins standardises permettent a differents logiciels de portefeuille de reconstruire le meme ensemble de cles a partir de la meme graine.
