## Deux transactions pour des milliers de paiements

Un canal de paiement est un contrat multisig 2-of-2 sur la blockchain Bitcoin. Deux parties verrouillent des fonds dans une sortie partagee. Elles echangent des transactions d'engagement signees hors chaine pour mettre a jour la repartition du solde. Seules deux transactions sur la chaine sont necessaires : une pour ouvrir le canal et une pour le fermer.

Entre l'ouverture et la fermeture, les parties peuvent echanger des milliers de paiements. Chaque paiement met a jour la transaction d'engagement qui repartit les fonds verrouilles. La blockchain ne voit jamais ces etats intermediaires. Le resultat : des paiements rapides (millisecondes), des frais quasi nuls et aucun delai de confirmation de bloc.
