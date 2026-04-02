## Ark : UTXO hors chaine

Ark est un protocole de couche 2 qui permet des transactions Bitcoin bon marche sans necessiter de canaux. Un Ark Service Provider (ASP) coordonne des tours ou les utilisateurs echangent des Virtual UTXO (VTXO) : des transactions Bitcoin valides gardees hors chaine.

Chaque VTXO est une feuille dans un arbre de transactions enracine dans un seul UTXO partage sur la chaine. Les utilisateurs peuvent toujours diffuser leur VTXO pour reclamer des fonds sur la chaine (sortie unilaterale). En fonctionnement normal, l'ASP regroupe des milliers de transferts en un seul reglement sur la chaine.

Ark ne necessite pas l'ouverture de canaux ni le verrouillage de liquidite a l'avance. Les utilisateurs recoivent des VTXO en cedant les anciens lors de tours periodiques. Des swaps atomiques via des sorties connecteur assurent qu'aucun fonds n'est perdu pendant l'echange.
