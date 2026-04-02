## Scripts de verrouillage

Chaque sortie de transaction contient un script de verrouillage, appele scriptPubKey. Ce script definit les conditions a remplir pour depenser la sortie.

Le script de verrouillage le plus facile a comprendre est Pay-to-Public-Key-Hash (P2PKH). Il dit : "Pour depenser cette sortie, prouvez que vous possedez la cle privee correspondant a ce hachage de cle publique." Les transactions modernes utilisent des formats plus recents (P2WPKH pour SegWit, P2TR pour Taproot), mais P2PKH demontre le concept de base le plus clairement.

En notation de script, un verrouillage P2PKH ressemble a : OP_DUP OP_HASH160 <pubkey_hash> OP_EQUALVERIFY OP_CHECKSIG. Chaque element est une instruction que le moteur de script Bitcoin execute.
