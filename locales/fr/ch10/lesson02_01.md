## Trouver des pairs

Quand un noeud Bitcoin demarre pour la premiere fois, il ne connait aucun autre noeud. Il doit trouver au moins un pair pour commencer a participer au reseau.

Bitcoin utilise des **DNS seeds** pour la decouverte initiale. Les DNS seeds sont des serveurs DNS operes par des membres de la communaute Bitcoin. Ils renvoient les adresses IP de noeuds complets connus et stables. Bitcoin Core a plusieurs DNS seeds codes en dur dans son code source.

Si le DNS n'est pas disponible (bloque ou censure), Bitcoin Core inclut aussi une liste d'adresses IP codees en dur en dernier recours. Ces adresses sont mises a jour a chaque version du logiciel.

Une fois qu'un noeud se connecte a son premier pair, il peut lui demander des adresses supplementaires. Le noeud construit un ensemble diversifie de connexions, generalement 8 sortantes et jusqu'a 125 entrantes.
