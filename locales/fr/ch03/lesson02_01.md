## Ce que fait un noeud complet

Un noeud complet telecharge et verifie independamment chaque bloc de la blockchain, en partant du bloc de genese de 2009. Il verifie chaque transaction par rapport a chaque regle de consensus. Il maintient sa propre copie de l'ensemble UTXO (toutes les sorties de transaction non depensees).

Quand vous executez un noeud complet, vous ne faites confiance a personne. Votre noeud verifie l'historique complet du reseau par lui-meme. Si un bloc ou une transaction viole une regle, votre noeud le rejette, quel que soit le nombre d'autres noeuds qui l'acceptent.

Un noeud complet relaie egalement les transactions et blocs valides aux autres noeuds, contribuant a la sante et a la resilience du reseau.
