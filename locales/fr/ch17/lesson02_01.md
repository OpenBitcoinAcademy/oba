## Signature en deux tours

La signature FROST se deroule en deux tours entre les signataires participants. Un coordinateur relaie les messages mais n'apprend jamais assez pour forger une signature ou extraire des parts secretes. Le coordinateur n'est pas de confiance.

Au premier tour, chaque signataire genere une paire de nonces aleatoires et envoie les engagements de nonce correspondants au coordinateur. Le coordinateur collecte les engagements de tous les t signataires participants et les diffuse au groupe. Ces engagements lient chaque signataire a ses nonces avant que quiconque ne voie les valeurs des autres signataires, empechant un signataire malveillant de choisir ses nonces de maniere adaptative pour manipuler le resultat.

Au second tour, chaque signataire calcule le nonce agrege a partir des engagements collectes, construit le hachage de defi Schnorr et produit une part de signature en utilisant sa part de cle secrete et son nonce. Chaque signataire envoie sa part de signature au coordinateur.
