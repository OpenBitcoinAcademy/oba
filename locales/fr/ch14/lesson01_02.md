## Pourquoi le tweak fonctionne

Le tweak lie l'arbre de scripts a la cle publique de maniere cryptographique. Modifier un script dans l'arbre change la racine de Merkle, ce qui change le tweak, ce qui change Q. Un script engage dans Q au moment de la creation de la sortie ne peut pas etre modifie par la suite.

La cle privee est modifiee de maniere identique : tweaked_privkey = (privkey + t) mod n. Le detenteur de la cle privee interne peut calculer la cle privee modifiee et signer directement. C'est la depense par chemin de cle.

Les hachages etiquetes empechent les collisions entre differentes utilisations de la fonction de hachage. L'etiquette "TapTweak" est separee par domaine : SHA256(SHA256("TapTweak") || SHA256("TapTweak") || data). Differentes etiquettes produisent des sorties de hachage differentes pour les memes donnees, eliminant les attaques inter-protocoles.
