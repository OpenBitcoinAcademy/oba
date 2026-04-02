## Filtres de Bloom et leur probleme de vie privee

**BIP 37** a introduit les filtres de Bloom comme moyen pour les clients SPV de demander uniquement les transactions pertinentes aux noeuds complets. Un filtre de Bloom est une structure de donnees probabiliste. Le client cree un filtre contenant ses adresses et l'envoie a un pair. Le pair envoie uniquement les transactions correspondant au filtre.

La conception a une faille serieuse de vie privee. Le noeud complet voit le filtre de Bloom et peut deduire quelles adresses appartiennent au client. Meme si un filtre de Bloom a des faux positifs, la recherche a montre qu'un noeud malveillant peut identifier les adresses du client avec une grande precision en analysant le motif de bits du filtre.

Un noeud servant des requetes de filtres de Bloom supporte aussi un cout de calcul important. Il doit tester chaque transaction de chaque bloc par rapport au filtre. Cela cree un vecteur de deni de service : un attaquant peut se connecter plusieurs fois avec differents filtres et forcer le noeud a effectuer un travail couteux. De nombreux operateurs de noeuds desactivent maintenant le support de BIP 37.
