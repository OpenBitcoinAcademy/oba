## La chaine de hachages

Chaque en-tete de bloc contient le hachage de l'en-tete du bloc precedent. Le bloc 500 000 contient le hachage du bloc 499 999, qui contient le hachage du bloc 499 998, et ainsi de suite jusqu'au bloc de genese (bloc 0).

Modifier des donnees dans le bloc 499 999 changerait son hachage. Le bloc 500 000 contiendrait alors un hachage de bloc precedent incorrect et deviendrait invalide. Pour modifier un bloc historique, un attaquant doit refaire la preuve de travail de ce bloc et de chaque bloc qui suit.

Ce chainage cumulatif est ce qui rend la blockchain inviolable. Plus un bloc est profond, plus le travail necessaire pour le modifier est important. Chaque nouveau bloc ajoute une couche de protection supplementaire a chaque bloc situe en dessous.
