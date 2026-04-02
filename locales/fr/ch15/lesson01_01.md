## Ecrire du Bitcoin Script a la main

Bitcoin Script est un langage base sur une pile. Chaque condition de depense est une sequence d'opcodes : empiler une cle, verifier une signature, verifier un hachage, appliquer un verrou temporel. Pour une depense a cle unique, le script est court et bien compris.

Combinez des conditions et la difficulte augmente vite. Un multisig 2-of-3 avec un repli a verrou temporel necessite un ordonnancement soigneux des branches OP_IF, OP_CHECKMULTISIG et OP_CHECKSEQUENCEVERIFY. Un opcode mal place change entierement les conditions de depense. Une cle dupliquee pourrait passer inapercue jusqu'a ce que les fonds soient bloques.

Aucun outil ne peut analyser un Script arbitraire et repondre : "Qui peut depenser cela, et sous quelles conditions ?" Le langage manque de structure. Differentes sequences d'opcodes peuvent encoder la meme politique, et il n'y a pas de methode generale pour les comparer.
