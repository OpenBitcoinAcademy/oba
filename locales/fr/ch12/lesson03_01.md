## Quatre processus de consensus

Le consensus de Bitcoin n'est pas un mecanisme unique. C'est le resultat de quatre processus independants executés sur chaque noeud du reseau.

Premierement, chaque noeud verifie independamment chaque transaction par rapport a un ensemble de regles (signatures valides, entrees non depensees, montants corrects). Deuxiemement, les mineurs aggregent les transactions verifiees dans des blocs candidats. Troisiemement, chaque noeud valide independamment les nouveaux blocs par rapport aux regles de consensus (preuve de travail correcte, coinbase valide, structure correcte). Quatriemement, chaque noeud selectionne independamment la chaine avec le plus de preuve de travail cumulee.

Aucune etape seule n'est "le consensus". Les quatre processus, executés sur des milliers de noeuds independants, produisent un accord emergent. Aucun noeud ne fait confiance a un autre. Chaque noeud verifie tout par lui-meme.
