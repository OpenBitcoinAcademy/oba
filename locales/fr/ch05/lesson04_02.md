## La phrase de passe optionnelle

BIP39 prend en charge une phrase de passe optionnelle qui est melangee dans la derivation de la graine. Les memes 12 mots avec des phrases de passe differentes produisent des graines differentes et des portefeuilles differents.

Cela offre un deni plausible : un utilisateur peut creer un portefeuille leurre avec une phrase de passe et un vrai portefeuille avec une autre. Sous contrainte, il peut reveler la phrase de passe du leurre sans exposer les fonds principaux.

Le risque : oublier la phrase de passe signifie perdre l'acces au portefeuille. Il n'y a pas de mecanisme de recuperation. La phrase de passe n'est stockee nulle part. Elle n'existe que dans la memoire de l'utilisateur ou sa sauvegarde physique.
