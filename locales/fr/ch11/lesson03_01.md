## Qu'est-ce qu'un arbre de Merkle ?

Un arbre de Merkle est un arbre binaire de hachages. Chaque noeud feuille est le hachage d'une transaction. Chaque noeud interne est le hachage de ses deux enfants concatenes. La racine de l'arbre (la racine de Merkle) est un hachage unique qui engage sur chaque transaction du bloc.

Si un bloc contient quatre transactions (A, B, C, D), l'arbre est : hash(A) et hash(B) se combinent en hash(AB). hash(C) et hash(D) se combinent en hash(CD). hash(AB) et hash(CD) se combinent en la racine de Merkle.

Si le nombre de transactions est impair, le dernier hachage est duplique pour faire un compte pair a chaque niveau.
