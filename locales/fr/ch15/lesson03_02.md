## Descripteurs et Miniscript ensemble

Les descripteurs supportent l'integration d'expressions Miniscript. Le descripteur `wsh()` enveloppe une expression Miniscript dans une sortie P2WSH. Le descripteur `tr()` place une expression Miniscript dans une feuille d'arbre de script Taproot.

Un descripteur comme `wsh(and_v(v:pk(Alice),or_d(pk(Bob),older(12960))))` definit une politique de depense complete en une seule chaine. N'importe quel portefeuille qui comprend le format de descripteur peut importer cette chaine, deriver les adresses correctes, identifier les conditions de depense et construire des transactions valides. Le descripteur porte la politique complete, pas un indice partiel.

C'est ce qui rend les portefeuilles interoperables. Un signataire materiel peut afficher les conditions de depense extraites du descripteur. Un portefeuille en lecture seule peut surveiller les adresses resultantes. Un coordinateur de signature peut construire un PSBT qui satisfait la politique. Chaque outil lit la meme chaine de descripteur et arrive au meme Script de sortie.

Les descripteurs incluent une somme de controle : huit caracteres ajoutes apres un symbole `#`. La somme de controle detecte les erreurs de transcription. Un portefeuille rejette tout descripteur dont la somme de controle ne correspond pas. Un seul caractere modifie dans la cle ou la politique produit une somme de controle differente, detectant les erreurs avant que des fonds ne soient envoyes a la mauvaise adresse.
