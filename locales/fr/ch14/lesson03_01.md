## Quand le chemin de cle echoue

Si le chemin de cle ne peut pas etre utilise (un signataire est indisponible, un verrou temporel doit etre exerce), le chemin de script offre une alternative. Le depenseur revele un script de l'arbre engage dans Q, prouve qu'il a ete engage au moment de la creation, et satisfait ce script.

Le temoin pour une depense par chemin de script contient : les donnees qui satisfont le script (signatures, preimages), le script lui-meme et un bloc de controle. Le bloc de controle contient la cle interne P, un octet de version de feuille et la preuve de Merkle (les hachages des noeuds freres le long du chemin de la feuille de script a la racine).

Le verificateur reconstruit la racine de Merkle a partir du script revele et de la preuve, calcule le tweak attendu et verifie que Q est egal a P + tweak * G. Si les calculs sont corrects, le script a ete engage au moment de la creation.
