## Routage en oignon

Lightning enveloppe les instructions de paiement dans des couches de chiffrement, comme un oignon. Chaque noeud de transfert dechiffre une couche, qui revele uniquement le prochain saut et le montant a transmettre. Le noeud n'apprend ni l'expediteur, ni le destinataire final, ni la longueur totale du chemin.

Ce modele de vie privee (base sur SPHINX) signifie que Bob, transferant un paiement d'Alice a Carol, sait uniquement qu'Alice lui a envoye quelque chose et qu'il doit le transmettre a Carol. Il ne sait pas si Alice est l'expediteur original ou un autre noeud de transfert. Il ne sait pas si Carol est le destinataire final.

L'oignon a une taille fixe quel que soit le nombre de sauts, empechant l'analyse de longueur de chemin.
