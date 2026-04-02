## Correspondance bijective

La relation entre Miniscript et Script est bijective pour l'ensemble de fragments supporte. Chaque expression Miniscript correspond a exactement un encodage Script, et chaque motif Script supporte correspond a exactement une expression Miniscript.

Considerez le fragment `pk(K)` : il encode en `<K> OP_CHECKSIG`. Le fragment `and_v(v:pk(A),pk(B))` encode en `<A> OP_CHECKSIGVERIFY <B> OP_CHECKSIG`. Il n'y a pas d'ambiguite. A partir du Script, vous pouvez retrouver le Miniscript original.

Cette propriete rend l'analyse aller-retour possible. Un portefeuille peut recevoir un Script d'une autre partie, le decoder en Miniscript et determiner les conditions de depense sans faire confiance a la description de l'expediteur. Deux portefeuilles construits par des equipes differentes peuvent s'accorder sur une politique, compiler en Miniscript independamment et verifier qu'ils ont produit le meme Script.

Tous les Scripts Bitcoin valides ne peuvent pas etre representes en Miniscript. L'ensemble de fragments couvre les motifs necessaires pour les politiques de depense pratiques : cles, hachages, verrous temporels, seuils et combinaisons booleennes. Les scripts qui utilisent des sequences d'opcodes inhabituelles ou des manipulations de pile sortent du sous-ensemble Miniscript.
