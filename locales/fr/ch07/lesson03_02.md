## SegWit enveloppe (P2SH-P2WPKH)

Tous les portefeuilles ne prenaient pas en charge les adresses segwit natives (bc1...) quand segwit a ete active. Le segwit enveloppe fournit une retro-compatibilite en placant le programme segwit dans un script de rachat P2SH.

La sortie ressemble a une adresse P2SH standard (commencant par "3"). Le script de rachat contient : OP_0 <20-byte-pubkey-hash>. Lors de la depense, le script d'entree revele le script de rachat, et le temoin fournit la signature et la cle publique.

Le segwit enveloppe est un format de transition. Le segwit natif (adresses bc1q...) est prefere pour les nouvelles transactions car il est plus petit, moins cher et a une meilleure detection d'erreurs via l'encodage bech32.
