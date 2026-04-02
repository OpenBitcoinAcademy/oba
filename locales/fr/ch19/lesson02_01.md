## Single-Use Seals

RGB definit un "sceau" comme un UTXO Bitcoin. Quand une transition d'etat se produit (transfert d'un jeton, mise a jour d'un contrat), le sceau est "ferme" en depensant cet UTXO. Un UTXO ne peut etre depense qu'une seule fois, donc un sceau ne peut etre ferme qu'une seule fois. Cela empeche la double depense des actifs RGB en utilisant le mecanisme de consensus existant de Bitcoin.

Les donnees de transition d'etat (qui possede quoi, mises a jour de contrats) ne touchent jamais la blockchain. Seul un engagement cryptographique sur la transition est embarque dans la transaction Bitcoin, generalement dans une sortie Taproot. La blockchain ancre le timing et l'ordonnancement. Les donnees vivent hors chaine avec les parties impliquees.
