## Des graines aux arbres

Une graine est un nombre aleatoire, generalement de 128 a 256 bits. C'est la racine de toute derivation de cles. A partir de cette graine, une hierarchie de cles peut etre generee : la graine produit une cle maitresse, la cle maitresse produit des cles enfants, chaque enfant peut produire ses propres enfants.

Cette structure arborescente est la base des portefeuilles Hierarchiques Deterministes (HD), definis dans BIP32. L'arbre permet d'organiser les cles par fonction : une branche pour recevoir des paiements, une autre pour la monnaie, une autre pour un compte different.

La graine ne change jamais. Chaque cle de l'arbre peut etre regeneree a partir d'elle. Une seule sauvegarde protege un nombre illimite de futures cles.
