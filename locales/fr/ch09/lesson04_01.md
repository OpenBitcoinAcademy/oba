## Package Relay

Historiquement, Bitcoin Core evaluait chaque transaction independamment pour l'admission au mempool. Une transaction parent a frais bas etait rejetee meme si une transaction enfant a frais eleves en dependait.

Le package relay change cela. Les noeuds evaluent des groupes de transactions liees ensemble, acceptant un parent a frais bas si son enfant amene le taux de frais combine au-dessus du seuil.

Cela ameliore la fiabilite de CPFP. Sans package relay, le parent pourrait ne pas se propager aux mineurs, rendant l'enfant inutile. Avec le package relay, le parent et l'enfant voyagent ensemble a travers le reseau.
