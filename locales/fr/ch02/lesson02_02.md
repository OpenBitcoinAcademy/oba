## Propagation a travers le reseau

Le portefeuille d'Alice signe la transaction avec sa cle privee et la diffuse sur le reseau Bitcoin. Son portefeuille se connecte a plusieurs noeuds, qui verifient la transaction independamment : format correct, signature valide, entrees non depensees, valeurs des sorties ne depassant pas celles des entrees.

Chaque noeud qui accepte la transaction l'ajoute a son mempool (une liste locale de transactions non confirmees) et la relaie a ses pairs. En quelques secondes, la transaction atteint des milliers de noeuds a travers le reseau. Aucun serveur central ne coordonne cela. Chaque noeud suit les memes regles independamment.

Le portefeuille de Bob surveille le reseau pour les transactions payant a son adresse. Quand il voit la transaction d'Alice, elle apparait comme "non confirmee" dans le portefeuille de Bob. Le paiement a ete diffuse mais pas encore enregistre dans un bloc.
