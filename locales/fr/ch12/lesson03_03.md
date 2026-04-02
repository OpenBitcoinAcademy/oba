## Forks et reorganisations

Deux mineurs trouvent parfois des blocs valides presque en meme temps. Quand cela arrive, des parties du reseau voient un bloc en premier, et d'autres parties voient l'autre. La chaine se divise temporairement en deux branches concurrentes.

C'est un fork naturel, pas une attaque. Chaque noeud travaille sur la branche qu'il a recue en premier. L'egalite se brise quand le bloc suivant est trouve. Si un mineur prolonge une branche, cette branche a maintenant plus de preuve de travail cumulee. Les noeuds sur la branche plus courte basculent vers la plus longue. Ce basculement est appele une reorganisation. Les transactions dans le bloc abandonne retournent au mempool pour inclusion dans un futur bloc.

La plupart des forks naturels se resolvent en un bloc. Les reorganisations plus profondes sont rares et deviennent exponentiellement moins probables avec chaque confirmation supplementaire.
