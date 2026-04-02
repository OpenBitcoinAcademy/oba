## La vue d'ensemble

Les signatures a seuil representent un changement dans le fonctionnement de la garde Bitcoin. L'approche legacy revelait les politiques de depense sur la chaine : n'importe qui pouvait compter les cles, voir le seuil et suivre les portefeuilles multisig a travers les transactions. FROST et ChillDKG deplacent toute cette complexite hors chaine.

Le groupe negocie son seuil et genere les cles en utilisant ChillDKG. N'importe quels t signataires produisent une signature en utilisant le protocole FROST. La blockchain voit une depense par chemin de cle Taproot. Les verificateurs verifient une signature contre une cle. Le cout sur la chaine est constant quel que soit le nombre de participants impliques : 64 octets pour la signature, 32 octets pour la cle.

C'est possible parce que la verification Taproot et Schnorr BIP 340 sont deja deployes sur le reseau Bitcoin. Aucun soft fork n'est requis. Aucun nouvel opcode. Le protocole de signature a seuil s'execute entierement dans les portefeuilles des participants. La couche de consensus verifie le resultat en utilisant les memes regles qu'elle utilise pour chaque autre depense Taproot.
