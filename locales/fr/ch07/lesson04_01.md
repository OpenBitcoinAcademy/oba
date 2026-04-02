## Taproot (P2TR)

Taproot (BIP341) est le type de sortie le plus recent, active en novembre 2021 comme segwit version 1. Une sortie Taproot engage un commit sur une seule cle publique. La depense peut se faire de deux manieres.

**Chemin par cle** : le proprietaire signe directement avec la cle engagee. Sur la blockchain, cela semble identique a une depense a signature unique. Aucun script n'est revele. Personne ne peut dire si la sortie avait des conditions de depense supplementaires.

**Chemin par script** : le proprietaire revele un script d'un arbre de Merkle de scripts engages dans la sortie. Des conditions complexes sont possibles (multisig, timelocks, hash locks) tout en gardant le cas courant (chemin par cle) compact et prive.

Les adresses Taproot commencent par "bc1p" sur le mainnet et utilisent l'encodage bech32m.
