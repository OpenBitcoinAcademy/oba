## Preuves de Merkle

Une preuve de Merkle permet de verifier qu'une transaction specifique est incluse dans un bloc sans telecharger chaque transaction. La preuve consiste en le hachage de la transaction, les hachages des noeuds freres le long du chemin vers la racine, et la racine de Merkle de l'en-tete du bloc.

Pour verifier, commencez avec le hachage de la transaction. Hachez-le avec son frere. Hachez le resultat avec le frere suivant. Continuez jusqu'a atteindre la racine. Si votre racine calculee correspond a la racine de Merkle dans l'en-tete du bloc, la transaction est dans le bloc.

Un bloc avec 4 000 transactions necessite une preuve de Merkle d'environ 12 hachages (log2(4000)). Les clients SPV utilisent les preuves de Merkle pour verifier les transactions sans telecharger les blocs complets.
