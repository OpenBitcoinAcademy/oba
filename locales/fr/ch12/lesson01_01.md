## Le minage comme loterie decentralisee

Bitcoin n'a pas d'autorite centrale qui decide quelles transactions entrent dans les blocs. A la place, les mineurs rivalisent dans une loterie computationnelle. Le gagnant gagne le droit d'ajouter le prochain bloc a la chaine.

Chaque mineur prend un ensemble de transactions non confirmees, les assemble dans un bloc candidat et hache l'en-tete du bloc avec SHA-256 (applique deux fois). Le resultat doit etre un nombre inferieur a la cible actuelle. La plupart des tentatives echouent. Le mineur change le champ nonce dans l'en-tete et hache a nouveau.

Ce processus se repete des milliards de fois par seconde a travers tous les mineurs mondiaux. Le mineur qui trouve un hachage valide en premier diffuse le bloc. Les autres noeuds le verifient independamment. Personne n'a besoin de permission pour miner, et personne ne peut predire qui gagnera la prochaine fois.
