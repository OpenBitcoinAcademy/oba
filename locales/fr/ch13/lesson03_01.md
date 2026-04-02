## Construire des systemes Bitcoin securises

Les developpeurs qui construisent sur Bitcoin font face a un defi unique : les bogues peuvent perdre de l'argent. Un defaut dans la generation de cles, la construction de transactions ou la validation de signatures peut entrainer le vol ou la perte permanente de fonds.

Le modele de consensus decentralise signifie qu'il n'y a pas d'autorite pour annuler une transaction erronee. Le code qui manipule des cles privees doit les traiter comme les donnees les plus sensibles du systeme. Les cles doivent etre generees a partir de sources d'entropie de haute qualite, stockees chiffrees au repos et effacees de la memoire apres utilisation.

Chaque composant qui touche aux cles ou construit des transactions devrait etre audite, teste contre des vecteurs connus et soumis a un examen contradictoire.
