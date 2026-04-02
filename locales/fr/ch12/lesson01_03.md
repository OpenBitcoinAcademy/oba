## Reajustement de la difficulte

Bitcoin ajuste sa difficulte tous les 2 016 blocs, environ toutes les deux semaines. L'objectif est de maintenir un intervalle moyen de 10 minutes entre les blocs.

A chaque point de reajustement, les noeuds calculent combien de temps les 2 016 blocs precedents ont pris. Si les blocs sont arrives plus vite que 10 minutes en moyenne, la cible diminue (la difficulte augmente). Si les blocs sont arrives plus lentement, la cible augmente (la difficulte baisse). La formule compare le temps ecoule reel aux 20 160 minutes attendues.

Une securite empeche la difficulte de changer par plus d'un facteur de quatre en un seul reajustement. Cela limite la vitesse a laquelle la difficulte peut monter ou chuter. Le mecanisme est entierement algorithmique. Aucun comite ne vote dessus. Chaque noeud calcule la meme nouvelle cible a partir des memes donnees de chaine.
