## MAST et Tapscript

L'arbre de Merkle dans Taproot est un Merklized Alternative Script Tree (MAST). Chaque feuille de l'arbre est un script different (une condition de depense differente). Pour depenser via le chemin par script, le depenseur revele uniquement le script qu'il utilise et fournit une preuve de Merkle montrant qu'il fait partie de l'arbre engage.

Les scripts non utilises restent caches. Une sortie Taproot avec 100 conditions de depense possibles a la meme apparence sur la blockchain qu'une avec 1 condition, a moins que le chemin par script ne soit exerce. Les branches inutilisees restent privees.

Tapscript (BIP342) est le langage de script utilise dans les feuilles du chemin par script de Taproot. Il est similaire au Script legacy mais avec des ameliorations : OP_CHECKSIGADD remplace OP_CHECKMULTISIG (permettant la validation par lots), les opcodes de signature utilisent Schnorr au lieu d'ECDSA, et la limite de taille du script est supprimee (remplacee par des limites de poids).
