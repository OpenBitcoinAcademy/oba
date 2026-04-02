## Schnorr contre ECDSA

Les deux algorithmes utilisent la courbe secp256k1. Les deux fournissent une securite equivalente contre la recuperation de cle. Les differences sont pratiques.

Les signatures Schnorr sont plus petites (64 octets contre 71-73). Schnorr supporte l'agregation native de signatures. Schnorr a une equation de verification plus elegante. Schnorr a une preuve de securite formelle.

ECDSA a ete choisi pour Bitcoin en 2008 parce que Schnorr etait brevete jusqu'en 2008, et le statut du brevet etait incertain quand Satoshi a concu le systeme. Avec le brevet expire et Taproot active, les nouvelles transactions peuvent utiliser Schnorr. ECDSA reste supporte pour la retro-compatibilite.
