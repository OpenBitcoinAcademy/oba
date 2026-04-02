## Trois proprietes cles

Les fonctions de hachage ont trois proprietes qui les rendent utiles pour Bitcoin.

**Deterministe.** La meme entree produit toujours la meme sortie. Hachez "bitcoin" un million de fois et vous obtenez le meme resultat a chaque fois.

**Effet avalanche.** Changez un seul bit de l'entree et la sortie change completement. "bitcoin" et "Bitcoin" produisent des hachages totalement differents. Il n'y a aucun moyen de predire comment la sortie va changer.

**Sens unique.** Etant donne une sortie de hachage, il n'y a aucun moyen de remonter pour trouver l'entree. Vous pouvez aller de l'entree au hachage, mais pas du hachage a l'entree. La seule option est de deviner des entrees et de verifier.
