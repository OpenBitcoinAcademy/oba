## La poignee de main

Deux noeuds etablissent une connexion avec une poignee de main version/verack sur le port TCP 8333 (le port par defaut du reseau Bitcoin).

Le noeud qui se connecte envoie un message `version` contenant sa version de protocole, sa hauteur de bloc, l'heure actuelle et les services qu'il supporte. Le noeud recepteur repond avec son propre message `version`. Chaque noeud envoie ensuite un `verack` (accusé de reception de version) pour confirmer.

Apres la poignee de main, les noeuds peuvent echanger des donnees. Si les versions de protocole sont incompatibles, la connexion est interrompue.

Les noeuds partagent les adresses de pairs connues en utilisant les messages `addr` et `getaddr`. Quand un noeud apprend de nouvelles adresses, il les stocke et peut s'y connecter plus tard. Ce protocole de propagation permet au reseau de croitre et de se reparer sans aucun annuaire central.
