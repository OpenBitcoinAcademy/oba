## Agregation de multisignatures

Avec ECDSA, une transaction necessitant trois signataires a besoin de trois signatures separees, chacune de 71-73 octets. Les signatures sont verifiees individuellement.

Avec Schnorr, les memes trois signataires peuvent combiner leurs signatures en une seule signature agregee de 64 octets en utilisant des protocoles comme MuSig2. Sur la blockchain, l'agregat semble identique a une signature a signataire unique. Aucun observateur ne peut dire combien de parties ont participe.

Cela a deux avantages. Les transactions avec plusieurs signataires utilisent moins d'espace de bloc (frais plus bas). Et les transactions multisignatures deviennent indiscernables des transactions a signature unique, ameliorant la vie privee.
