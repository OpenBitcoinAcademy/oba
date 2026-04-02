## SHA-256 dans Bitcoin

Bitcoin utilise une fonction de hachage appelee SHA-256, concue par la National Security Agency des Etats-Unis. SHA signifie Secure Hash Algorithm. 256 fait reference a la taille de la sortie : 256 bits, soit 32 octets.

SHA-256 apparait partout dans Bitcoin. Les identifiants de transaction sont des hachages double-SHA-256. Le minage consiste a trouver des entrees qui produisent des hachages inferieurs a une valeur cible. La derivation d'adresse utilise SHA-256 combinee avec RIPEMD-160 (une fonction de hachage differente qui produit une sortie plus courte de 20 octets). La combinaison des deux, appelee HASH160, est ce qui cree le hachage de cle publique dans une adresse.
