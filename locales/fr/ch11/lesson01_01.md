## Qu'est-ce qu'un bloc ?

Un bloc est un lot de transactions regroupees et ajoutees a la blockchain. Chaque bloc contient un en-tete et une liste de transactions.

L'en-tete du bloc a six champs : version, hachage du bloc precedent, racine de Merkle (un hachage de toutes les transactions du bloc), horodatage, cible de difficulte et nonce. Le hachage du bloc precedent lie chaque bloc au precedent, formant une chaine.

Les mineurs rivalisent pour trouver une valeur de nonce qui fait tomber le hachage de l'en-tete du bloc en dessous de la cible de difficulte. Cela necessite des milliers de milliards de tentatives. Le premier mineur a trouver un nonce valide diffuse le bloc. Les autres noeuds le verifient et l'ajoutent a leur copie de la chaine.
