## Selection des UTXO

Construire une transaction commence par choisir quelles sorties depenser. Votre portefeuille scanne la blockchain a la recherche d'UTXO verrouilles a vos adresses. Ce sont les pieces que vous controlez.

Pour envoyer 0.5 BTC, votre portefeuille selectionne un ou plusieurs UTXO dont la somme atteint au moins 0.5 BTC. Si le plus petit UTXO est de 0.8 BTC, cet UTXO unique devient l'entree. L'excedent (moins les frais) vous revient comme monnaie.

Si aucun UTXO n'est assez grand, le portefeuille combine plusieurs UTXO comme entrees separees dans la meme transaction. Chaque entree necessite sa propre signature.
