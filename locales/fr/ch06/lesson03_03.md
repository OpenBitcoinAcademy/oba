## Varints et endianness

Bitcoin utilise deux conventions d'encodage dans toute sa serialisation.

Les **varints** (entiers a longueur variable) encodent les compteurs et les longueurs. Une valeur inferieure a 253 tient dans un octet. Les valeurs 253-65535 utilisent un prefixe 0xFD suivi de 2 octets. Les valeurs jusqu'a environ 4 milliards utilisent 0xFE suivi de 4 octets. Les valeurs plus grandes utilisent 0xFF suivi de 8 octets.

L'ordre d'octets **little-endian** place l'octet le moins significatif en premier. La version 1 est serialisee comme 01 00 00 00, pas 00 00 00 01. Les valeurs en satoshis, les hauteurs de bloc et la plupart des champs numeriques utilisent le little-endian.

Les hachages de transaction (txids) sont affiches en ordre d'octets inverse par convention. Le hachage double-SHA-256 produit des octets dans un ordre, mais les explorateurs de blocs et Bitcoin Core les affichent inverses. C'est une convention d'affichage, pas une regle de serialisation.
