## Decrements de verrou temporel

Chaque saut dans la chaine a un verrou temporel plus court que le saut precedent. Alice donne 40 blocs a Bob. Bob donne 30 blocs a Carol. Ce decrement (appele le delta CLTV) assure que chaque noeud de transfert a le temps de reclamer les fonds de son saut en aval avant que son HTLC en amont n'expire.

Si Carol ne revele pas R dans les 30 blocs, le HTLC de Bob expire et les fonds retournent a Bob. Bob a encore 10 blocs pour regler avec Alice. Si Bob ne revele pas R a Alice dans les 40 blocs, les fonds d'Alice lui reviennent.

Le decrement empeche une attaque temporelle ou un noeud en aval retarde la revelation de la preimage jusqu'a ce que le HTLC en amont expire, piegeant le noeud de transfert.
