## Routage a la source

L'expediteur, pas le reseau, choisit le chemin de paiement. Chaque noeud Lightning maintient une vue locale de la topologie du reseau (quels noeuds existent, quels canaux les connectent, leurs capacites et taux de frais). L'expediteur calcule une route a partir de ces informations.

Cela differe du routage internet, ou chaque routeur decide independamment du prochain saut. Dans Lightning, l'expediteur a le controle total du chemin. Les noeuds intermediaires suivent les instructions de transfert sans connaitre la route complete.
