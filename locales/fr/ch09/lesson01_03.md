## Transactions conflictuelles

Alice peut creer deux transactions qui depensent la meme sortie : une payant Bob et une payant Carol. Les deux sont des signatures valides, mais une seule peut etre incluse dans la blockchain (la sortie ne peut etre depensee qu'une fois).

Les mineurs decident quelle transaction conflictuelle inclure. Par defaut, la plupart des mineurs suivent une politique "premier vu" et incluent la transaction qu'ils ont recue en premier. Mais c'est une politique, pas une regle de consensus. Un mineur peut inclure n'importe quelle transaction valide.

C'est pourquoi Bob devrait attendre des confirmations avant de considerer un paiement comme definitif. Les frais qu'Alice paie incitent les mineurs a inclure sa transaction rapidement, reduisant la fenetre pour les transactions conflictuelles.
