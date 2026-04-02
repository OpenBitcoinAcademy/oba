## BIP32 : le standard des portefeuilles HD

BIP32 definit comment deriver un arbre de cles a partir d'une seule graine. Le processus utilise HMAC-SHA512 pour diviser chaque etape de derivation en une cle enfant et un code de chaine. Le code de chaine est melange dans la derivation suivante, empechant quiconque de calculer les cles soeurs.

La cle maitresse se trouve a la racine de l'arbre. A partir d'elle, une sequence de cles enfants est derivee, chacune identifiee par un numero d'index (0, 1, 2, ...). Chaque enfant peut produire ses propres enfants, formant un arbre de profondeur arbitraire.

Cette structure arborescente donne aux portefeuilles un pouvoir d'organisation. Differentes branches de l'arbre servent des objectifs differents, et l'arbre entier est recuperable a partir de la graine originale.
