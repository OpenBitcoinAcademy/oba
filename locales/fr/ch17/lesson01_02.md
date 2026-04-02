## De n-of-n a t-of-n

MuSig2 a une limitation : il necessite que tous les n participants signent. Chaque detenteur de cle doit etre en ligne et cooperant au moment de la signature. Si un participant perd sa cle ou se deconnecte, les fonds sont bloques. Cela fait de MuSig2 un schema purement n-of-n.

FROST (Flexible Round-Optimized Schnorr Threshold signatures) resout cela. FROST est un schema de signature a seuil : n'importe quels t parmi n participants peuvent produire une signature valide. Une configuration FROST 3-of-5 signifie que trois des cinq detenteurs de cles peuvent signer. Les deux autres peuvent etre hors ligne, indisponibles ou meme perdus de maniere permanente.

La signature que FROST produit est une signature Schnorr standard de 64 octets. Sur la chaine, elle est indiscernable d'une depense a cle unique. La politique de seuil, le nombre de signataires et les cles publiques individuelles sont tous caches. Une depense FROST ressemble exactement a une depense par chemin de cle d'une seule adresse Taproot.
