## Combineur, Finaliseur, Extracteur

Le Combineur fusionne plusieurs PSBT contenant des signatures partielles pour la meme transaction. Dans un multisig 2-of-3, le signataire A produit un PSBT avec sa signature et le signataire B en produit un autre. Le Combineur prend les deux, fusionne les entrees PARTIAL_SIG pour chaque entree et produit un seul PSBT avec toutes les signatures disponibles.

Le Finaliseur transforme les signatures partielles en scriptSig et temoin complets pour chaque entree. Pour une entree P2PKH, il prend la seule PARTIAL_SIG et construit le scriptSig. Pour un multisig P2WSH 2-of-3, il prend les signatures partielles, les ordonne correctement et construit la pile temoin avec le script de rachat. Apres finalisation, les cartes d'entrees du PSBT contiennent les champs FINAL_SCRIPTSIG et FINAL_SCRIPTWITNESS. Les donnees partielles ne sont plus necessaires.

L'Extracteur lit le PSBT finalise et produit la transaction reseau brute. Il prend la transaction non signee de la carte globale, insere le FINAL_SCRIPTSIG et FINAL_SCRIPTWITNESS de chaque entree, et serialise le resultat. La sortie est une transaction Bitcoin standard prete a etre diffusee. Le PSBT a rempli son role.
