## SegWit natif (P2WPKH et P2WSH)

Segregated Witness a deplace les donnees de signature hors du script d'entree vers une structure temoin separee. Cela a corrige la malleabilite des transactions (des tiers pouvaient modifier le txid en alterant l'encodage de la signature) et a permis la ristourne temoin pour les frais.

P2WPKH (Pay to Witness Public Key Hash) est l'equivalent segwit de P2PKH. Le script de sortie contient : OP_0 <20-byte-pubkey-hash>. Le temoin fournit la signature et la cle publique. Le script d'entree est vide.

P2WSH (Pay to Witness Script Hash) est l'equivalent segwit de P2SH. Le script de sortie contient : OP_0 <32-byte-script-hash>. Le temoin fournit le script et ses donnees requises. P2WSH utilise un hachage SHA-256 de 32 octets au lieu du HASH160 de 20 octets de P2SH, offrant une resistance aux collisions plus forte.
