## Unites de poids

SegWit a introduit des unites de poids pour remplacer la metrique de taille basee sur les octets. Chaque octet de donnees non-temoin coute 4 unites de poids. Chaque octet de donnees temoin (signatures) coute 1 unite de poids.

Cette ristourne incite a utiliser les transactions segwit, qui stockent les signatures dans la structure temoin. Une transaction segwit utilise moins de poids qu'une transaction legacy avec le meme nombre d'entrees et de sorties.

Les octets virtuels (vbytes) convertissent le poids en equivalent octets : vbytes = poids / 4. Une transaction de 600 unites de poids fait 150 vbytes. Les taux de frais exprimes en sat/vB prennent en compte la ristourne segwit automatiquement.
