## Le probleme avant les PSBT

Une transaction Bitcoin a besoin d'informations provenant de plusieurs sources avant de pouvoir etre diffusee. Le portefeuille qui cree la transaction doit savoir quels UTXO depenser. Le signataire doit detenir les cles privees. Dans de nombreuses configurations, ce sont des appareils ou des logiciels differents.

Avant BIP 174, il n'existait pas de methode standard pour transmettre une transaction non signee entre ces participants. Chaque logiciel de portefeuille inventait son propre format. Bitcoin Core serialisait les transactions partielles differemment d'Electrum, qui les serialisait differemment des portefeuilles materiels. Une transaction creee dans un outil ne pouvait pas etre signee par un autre sans code de liaison personnalise.

Cela rendait les configurations multisignatures penibles. Chaque signataire avait besoin d'un logiciel compatible. Coordonner entre un ordinateur portable, un portefeuille materiel et une machine de stockage a froid necessitait des etapes manuelles et des conversions de format qui introduisaient des erreurs.
