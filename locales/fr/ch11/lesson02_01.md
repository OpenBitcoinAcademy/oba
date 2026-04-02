## Structure de l'en-tete de bloc

Chaque bloc commence par un en-tete de 80 octets contenant six champs.

**Version** (4 octets) : indique quelles regles de consensus le bloc suit et signale le support pour les propositions de soft fork via les versionbits.

**Hachage du bloc precedent** (32 octets) : le hachage double-SHA-256 de l'en-tete du bloc precedent. Cela lie chaque bloc a son predecesseur, formant la chaine.

**Racine de Merkle** (32 octets) : un hachage unique qui engage sur chaque transaction du bloc. Modifier n'importe quelle transaction change la racine de Merkle.

**Horodatage** (4 octets) : l'heure approximative a laquelle le bloc a ete mine, en secondes d'epoque Unix.

**Target Bits** (4 octets) : un encodage compact de la cible de preuve de travail. Le hachage de l'en-tete du bloc doit etre inferieur a cette cible.

**Nonce** (4 octets) : la valeur que les mineurs modifient pour chercher un hachage valide.
