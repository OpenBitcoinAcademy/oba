## La loterie du minage

Le minage est une loterie decentralisee. Chaque mineur cree un bloc candidat a partir des transactions de son mempool. Il hache l'en-tete du bloc de maniere repetee, en changeant un nombre appele nonce a chaque fois, a la recherche d'une valeur de hachage inferieure a une cible fixee par le reseau.

Trouver un hachage valide necessite des milliers de milliards de tentatives. Verifier un hachage valide ne necessite qu'un seul calcul. Cette asymetrie est le coeur de la preuve de travail : creer un bloc valide est couteux, en verifier un est bon marche.

Le mineur qui trouve un hachage valide en premier diffuse le bloc. Les autres mineurs le verifient, l'acceptent et commencent immediatement a travailler sur le bloc suivant. Le gagnant collecte la recompense de bloc : des bitcoins nouvellement crees (la subvention) plus la somme de tous les frais de transaction du bloc.
