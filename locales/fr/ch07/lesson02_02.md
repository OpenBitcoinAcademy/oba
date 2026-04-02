## Scripts multisignatures

Un script multisignature (multisig) necessite M signatures parmi N cles publiques pour autoriser une depense. Un multisig 2-of-3 necessite 2 signatures parmi 3 cles designees.

Dans une sortie multisig brute, le script de verrouillage contient : M <pubkey1> <pubkey2> ... <pubkeyN> N OP_CHECKMULTISIG. Le script de deverrouillage fournit : OP_0 <sig1> <sig2>. Le OP_0 initial est un contournement pour un bogue historique de decalage d'un dans OP_CHECKMULTISIG.

En pratique, le multisig est enveloppe dans P2SH. L'expediteur voit une adresse P2SH standard. Les details du multisig sont caches dans le script de rachat et reveles uniquement lors de la depense. Cela garde les sorties compactes et la complexite privee jusqu'au moment de la depense.
