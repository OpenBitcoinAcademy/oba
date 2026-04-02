## Cles publiques X-Only

Taproot utilise des cles publiques X-only de 32 octets au lieu du format compresse traditionnel de 33 octets. La coordonnee Y est implicitement choisie comme paire. Cela economise un octet par cle sur la chaine et dans les signatures.

Une cle publique compressee standard a un octet de prefixe (0x02 pour Y pair, 0x03 pour Y impair) suivi de 32 octets de coordonnee X. Taproot supprime le prefixe et exige que la coordonnee Y soit paire. Si le Q calcule a un Y impair, l'implementation inverse la cle privee (ce qui bascule Y en pair) avant de signer.

Cette convention simplifie la verification par lots et l'agregation de cles. Chaque cle de sortie Taproot fait 32 octets. Chaque signature Schnorr fait 64 octets. Il n'y a pas d'encodages a longueur variable.
