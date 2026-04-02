## Noeuds d'archive et noeuds elagues

Un noeud complet qui conserve chaque bloc depuis le bloc de genese est appele un **noeud d'archive**. Il peut fournir des donnees de blocs historiques a tout pair qui les demande. Les nouveaux noeuds rejoignant le reseau comptent sur les noeuds d'archive pour telecharger la blockchain complete.

Faire tourner un noeud d'archive necessite des centaines de gigaoctets d'espace disque. En 2024, la blockchain depasse 600 Go. Tous les operateurs de noeuds complets ne peuvent pas se permettre ce stockage.

Un **noeud elague** valide chaque bloc mais supprime les anciennes donnees de blocs apres les avoir traitees. Il ne conserve que l'ensemble UTXO et les blocs les plus recents. Un noeud elague applique chaque regle de consensus. Il ne repose pas sur la confiance. Le compromis : il ne peut pas fournir de blocs historiques a d'autres pairs.

Les noeuds d'archive et les noeuds elagues sont tous deux des noeuds complets. Les deux valident tout. La difference est de savoir s'ils stockent les anciens blocs pour que d'autres les telechargent.
