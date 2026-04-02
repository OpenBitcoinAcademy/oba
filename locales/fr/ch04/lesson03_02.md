## Les etapes de derivation

Creer une adresse Bitcoin a partir d'une cle publique necessite plusieurs etapes de hachage.

D'abord, hachez la cle publique avec SHA-256. Puis hachez ce resultat avec RIPEMD-160. Cela produit un hachage de 20 octets appele le hachage de cle publique, ou HASH160.

Ensuite, ajoutez un octet de version au debut. Cet octet identifie le reseau (mainnet ou testnet) et le type d'adresse.

Enfin, encodez le resultat. Les adresses legacy utilisent l'encodage Base58Check, qui ajoute une somme de controle de 4 octets et encode dans un alphabet lisible qui evite les caracteres confusants comme 0/O et l/1. Les adresses SegWit plus recentes utilisent l'encodage Bech32, qui est en minuscules uniquement et a une meilleure detection d'erreurs.
