## Du mempool a la blockchain

La transaction d'Alice se trouve dans le mempool, en attente qu'un mineur l'inclue dans un bloc. Les mineurs selectionnent des transactions de leur mempool, en privilegiant celles avec des taux de frais plus eleves par unite de poids.

Environ cinq minutes plus tard, un mineur trouve une solution de preuve de travail valide pour un nouveau bloc incluant la transaction d'Alice. Le mineur diffuse le bloc. Chaque noeud complet verifie le bloc independamment : preuve de travail valide, transactions valides, format correct. Les noeuds qui acceptent le bloc l'ajoutent a leur copie de la blockchain.

La transaction d'Alice a maintenant une confirmation. Le portefeuille de Bob met a jour le statut. Chaque bloc suivant construit au-dessus de celui-ci ajoute une confirmation supplementaire. Apres six confirmations (environ une heure), la transaction est consideree comme hautement securisee contre toute annulation.
