## Les trois cartes

Un PSBT contient trois types de cartes cle-valeur : une carte globale, une carte par entree et une carte par sortie. Chaque carte stocke des paires cle-valeur typees terminees par un octet 0x00. Le type de cle determine la signification de l'entree. La charge utile de cle et la charge utile de valeur portent les donnees.

La carte globale contient les donnees qui s'appliquent a la transaction entiere. Dans PSBTv0, l'entree globale la plus importante est UNSIGNED_TX (type de cle 0x00), qui contient la transaction non signee complete. Dans PSBTv2, la carte globale contient a la place TX_VERSION, FALLBACK_LOCKTIME, INPUT_COUNT et OUTPUT_COUNT comme champs separes. La transaction non signee est reconstruite a partir des cartes par entree et par sortie plutot que stockee en un seul bloc.

Les entrees globales incluent aussi XPUB (type de cle 0x01), qui fait correspondre une cle publique etendue a son chemin de derivation BIP 32. Cela permet a un signataire de verifier que les sorties de monnaie derivent de la meme racine de portefeuille sans avoir besoin d'acceder au descripteur complet du portefeuille.
