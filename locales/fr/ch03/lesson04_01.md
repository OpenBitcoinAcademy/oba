## Pourquoi les implementations multiples comptent

Un protocole defini par une seule implementation est fragile. Si Bitcoin Core a un bogue, chaque noeud qui l'execute est affecte simultanement. Un bogue dans le code de consensus pourrait diviser le reseau ou permettre des transactions invalides.

Plusieurs implementations independantes du meme protocole renforcent l'ecosysteme. Quand deux implementations divergent sur la validite d'un bloc, le desaccord revele un bogue dans l'une d'elles. C'est arrive en 2013 quand une difference entre BerkeleyDB et LevelDB a cause une scission de chaine involontaire. L'incident a demontre a la fois le risque de la monoculture et la valeur de la diversite.

Un protocole sain est un protocole que plusieurs equipes peuvent implementer a partir de la specification. Plus les implementations s'accordent sur chaque cas limite, plus le reseau a confiance dans l'exactitude et la bonne definition des regles.
