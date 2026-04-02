## Validation cote client

Dans RGB, seules les parties a un contrat valident ses transitions d'etat. Un emetteur de jetons et les detenteurs de jetons verifient l'historique des transferts. Aucun autre noeud du reseau n'a besoin de traiter ou de stocker ces donnees.

Comparez avec les systemes ou chaque noeud valide chaque contrat (le modele d'Ethereum). La validation cote client s'adapte sans limite : ajouter plus de contrats n'augmente pas la charge sur les noeuds non impliques. La vie privee est inherente : les donnees de contrat ne sont visibles que par les participants.

Le cout : les participants doivent stocker et verifier l'historique complet de leurs actifs. Si l'historique est perdu, l'actif ne peut pas etre verifie. RGB utilise des chaines d'engagement et des preuves pour rendre cette verification efficace.
