## Regles de validation des blocs

Quand un noeud recoit un nouveau bloc, il verifie des dizaines de criteres avant de l'accepter. Le hachage de l'en-tete du bloc doit etre inferieur a la cible actuelle. L'horodatage doit etre dans des limites acceptables. La premiere transaction doit etre un coinbase, et la sortie coinbase ne doit pas depasser la subvention autorisee plus les frais. Chaque autre transaction doit etre valide selon les regles de script.

La racine de Merkle du bloc doit correspondre a l'arbre de hachage de toutes les transactions incluses. La taille du bloc ne doit pas depasser la limite de consensus. Le bloc doit referencer un bloc parent valide que le noeud possede deja.

Si une verification echoue, le noeud rejette le bloc et peut se deconnecter du pair qui l'a envoye. Il n'y a pas de processus d'appel. Un bloc est soit valide, soit invalide. Cette validation stricte est ce qui empeche les mineurs de creer des bitcoins a partir de rien ou de depenser des pieces qu'ils ne possedent pas.
