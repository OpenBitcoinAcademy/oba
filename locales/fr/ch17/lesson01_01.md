## Le paysage du multisig

Bitcoin supporte la garde multi-parties depuis 2012 via OP_CHECKMULTISIG. Un script multisig 2-of-3 liste les trois cles publiques sur la chaine et exige deux signatures valides. Cela fonctionne, mais revele le nombre exact de parties qui controlent les fonds. Quiconque inspecte la blockchain voit le seuil, le nombre total de signataires et chaque cle publique impliquee.

Le cout sur la chaine augmente lineairement. Un multisig 3-of-5 place cinq cles publiques de 33 octets et trois signatures de 72 octets dans le temoin. Cela fait 381 octets de donnees temoin avant de compter le script lui-meme. Une configuration 7-of-10 coute encore plus. Les frais augmentent, la vie privee diminue, et la politique de depense est visible par tous les observateurs.

Taproot et les signatures Schnorr ont change ce qui est possible. MuSig2 permet a n participants de produire une seule signature agregee qui semble identique a une signature Schnorr solo sur la chaine. La cle publique combinee fait 32 octets. La signature fait 64 octets. Aucun observateur ne peut dire combien de parties etaient impliquees.
