## Construire l'arbre de scripts

L'arbre de Merkle est construit a partir de hachages TapLeaf et TapBranch.

Chaque feuille : tagged_hash("TapLeaf", leaf_version || compact_size(script_length) || script). La version de feuille pour le Tapscript actuel est 0xC0.

Chaque branche : tagged_hash("TapBranch", sorted(left_hash, right_hash)). Les deux enfants sont tries lexicographiquement avant le hachage, assurant que l'arbre a une seule forme canonique quel que soit l'ordre d'insertion.

L'arbre peut avoir jusqu'a 128 niveaux de profondeur, permettant jusqu'a 2^128 feuilles. Les arbres pratiques sont bien plus petits. Les scripts places a des profondeurs faibles (plus proches de la racine) ont des preuves de Merkle plus courtes et coutent moins cher a depenser. Placez le repli le plus probable a la profondeur 1. Placez les chemins d'urgence rarement utilises plus profondement.
