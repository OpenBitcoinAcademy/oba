## Limites de ressources et echecs caches

Le consensus Bitcoin impose des limites de ressources sur les scripts : une taille maximale de 10 000 octets en legacy, une limite de 201 opcodes non-push et un budget de sigops. Un script qui depasse une limite est invalide. La transaction qui le porte ne sera jamais confirmee.

Quand on compose des scripts a la main, la consommation de ressources est difficile a predire. Chaque branche d'un arbre OP_IF contribue differemment au compteur d'opcodes. Les conditions imbriquees multiplient la complexite. Un script peut fonctionner dans une branche et depasser la limite d'opcodes dans une autre, selon le chemin d'execution qu'un depenseur emprunte.

Ces echecs sont silencieux. Le script parait correct dans un editeur de texte. Il peut passer les tests unitaires pour un chemin de depense. Le chemin defaillant n'est decouvert que quand quelqu'un essaie de l'utiliser sur le reseau, et la transaction est rejetee.
