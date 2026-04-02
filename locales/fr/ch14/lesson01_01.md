## La cle interne et le tweak

Une sortie Taproot verrouille des fonds a une cle publique modifiee Q. Cette cle est derivee de deux entrees : une cle publique interne P et une racine de Merkle optionnelle d'un arbre de scripts.

La valeur du tweak t est calculee comme : t = tagged_hash("TapTweak", P || merkle_root). La cle modifiee est : Q = P + t * G, ou G est le point generateur sur secp256k1.

S'il n'y a pas d'arbre de scripts, le tweak utilise uniquement P : t = tagged_hash("TapTweak", P). La sortie engage toujours sur la cle interne, mais aucun script n'est embarque.

Sur la chaine, le scriptPubKey est : OP_1 suivi de la coordonnee X de 32 octets de Q. Cela fait 34 octets, la meme taille qu'une sortie P2WSH. Aucun observateur ne peut dire si Q contient un arbre de scripts embarque ou non.
