## L'algorithme ECDSA

ECDSA (Elliptic Curve Digital Signature Algorithm) etait le schema de signature original de Bitcoin. Il opere sur la courbe elliptique secp256k1, la meme courbe utilisee pour la generation de cles.

Pour signer un message, l'algorithme prend la cle privee et un hachage du message, les combine avec un nombre aleatoire (appele k, ou le nonce), et produit deux valeurs : r et s. Ensemble, (r, s) forment la signature.

Pour verifier, l'algorithme prend la cle publique, le hachage du message et la signature (r, s). Il effectue des calculs sur courbe elliptique pour verifier si la signature est coherente avec la cle publique et le message. Aucune cle privee n'est necessaire pour la verification.
