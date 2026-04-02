## Champs par entree

Chaque carte d'entree decrit une entree de la transaction. Les champs critiques dont un signataire a besoin sont les donnees UTXO, le chemin de derivation de la cle de signature et les scripts requis pour construire le temoin.

WITNESS_UTXO (type de cle 0x01) stocke la sortie depensee : le montant en satoshis et le scriptPubKey. Pour les entrees SegWit, cela suffit pour signer car l'algorithme de sighash engage directement sur le montant. NON_WITNESS_UTXO (type de cle 0x00) stocke la transaction precedente entiere, necessaire pour les entrees legacy ou le signataire doit verifier le montant en consultant la sortie dans la transaction complete.

BIP32_DERIVATION (type de cle 0x06) fait correspondre une cle publique a son chemin de derivation BIP 32 et l'empreinte de la cle maitresse. Le signataire fait correspondre l'empreinte a sa propre cle maitresse, derive la cle privee au chemin donne et signe. PARTIAL_SIG (type de cle 0x02) stocke une signature indexee par la cle publique qui l'a produite. SIGHASH_TYPE (type de cle 0x03) specifie quel drapeau sighash le signataire doit utiliser.

Pour les entrees P2SH, REDEEM_SCRIPT (type de cle 0x04) porte le script de rachat. Pour les entrees P2WSH, WITNESS_SCRIPT (type de cle 0x05) porte le script temoin. Le signataire en a besoin pour calculer le bon sighash.
