## Qu'est-ce qu'un PSBT

Une Partially Signed Bitcoin Transaction (PSBT) est un format binaire defini par BIP 174. Il enveloppe une transaction non signee avec les metadonnees dont chaque participant a besoin pour faire son travail. Les createurs attachent les donnees UTXO. Les metteurs a jour ajoutent les chemins de derivation. Les signataires ajoutent les signatures. Le format transporte tout le necessaire pour que chaque role puisse agir independamment.

Le binaire commence par un nombre magique de cinq octets : `0x70736274FF`. Les quatre premiers octets epelent "psbt" en ASCII. Le cinquieme octet, `0xFF`, marque le separateur. N'importe quel outil recevant un blob peut verifier ces cinq octets pour confirmer qu'il s'agit d'un PSBT avant de l'analyser.

Apres le nombre magique vient une sequence de cartes cle-valeur. Chaque entree de carte a un type de cle, une charge utile de cle et une charge utile de valeur. Un octet zero (0x00) termine chaque carte. La premiere carte est la carte globale. Les cartes suivantes alternent entre les cartes par entree et par sortie, une pour chaque entree et sortie de la transaction.
