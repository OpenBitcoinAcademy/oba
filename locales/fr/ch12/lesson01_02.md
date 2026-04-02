## Cible, difficulte et nonce

La cible est un nombre de 256 bits. Un hachage de bloc valide doit etre numeriquement inferieur a la cible. Une cible plus basse signifie que moins de hachages valides existent, ce qui rend la recherche plus difficile. La difficulte est l'inverse de la cible : quand la cible baisse, la difficulte augmente.

L'en-tete du bloc contient un champ nonce de 32 bits. Les mineurs incrementent cette valeur a chaque tentative de hachage, parcourant les $2^{32}$ (environ 4,3 milliards) possibilites. Les mineurs modernes epuisent l'espace du nonce en moins d'une seconde.

Quand l'espace du nonce est epuise, les mineurs modifient d'autres champs pour creer de nouvelles entrees de hachage. La technique la plus courante modifie le champ extra nonce de la transaction coinbase. Cela change la racine de Merkle dans l'en-tete du bloc, donnant au mineur un nouvel ensemble de $2^{32}$ nonces a essayer.
