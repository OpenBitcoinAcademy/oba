## L'exigence d'aleatoire

ECDSA et Schnorr necessitent tous deux un nombre aleatoire (nonce) lors de la creation de la signature. La securite du systeme entier depend de cette aleatoire.

Si le meme nonce est utilise pour signer deux messages differents avec la meme cle privee, la cle privee peut etre calculee a partir des deux signatures. Ce n'est pas un risque theorique. En 2013, une faille dans le generateur de nombres aleatoires d'Android a cause la reutilisation de nonces par plusieurs portefeuilles Bitcoin, et des fonds ont ete voles.

Les implementations modernes utilisent des nonces deterministes (RFC 6979 pour ECDSA, BIP340 pour Schnorr), derivant le nonce de la cle privee et du hachage du message. Cela elimine la dependance a un generateur de nombres aleatoires et empeche la reutilisation de nonces.
