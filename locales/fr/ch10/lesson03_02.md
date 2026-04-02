## Relais de blocs compacts

Defini dans **BIP 152**, le relais de blocs compacts reduit la bande passante necessaire pour propager un nouveau bloc. L'observation cle : la plupart des transactions dans un nouveau bloc sont deja dans le mempool du noeud recepteur.

Au lieu d'envoyer le bloc complet, le noeud annonceur envoie un message `cmpctblock`. Celui-ci contient l'en-tete du bloc, une liste d'identifiants de transaction courts et la transaction coinbase. Les identifiants courts sont des hachages tronques de 6 octets qui permettent au recepteur de faire correspondre les transactions qu'il a deja.

Le noeud recepteur reconstruit le bloc a partir de son propre mempool en utilisant les identifiants courts. Si des transactions manquent, il les demande avec un message `getblocktxn` et les recoit dans une reponse `blocktxn`.

BIP 152 definit deux modes. En **mode faible bande passante**, un noeud annonce d'abord le bloc avec un message `inv`. Le pair demande le bloc compact uniquement s'il est interesse. En **mode haute bande passante**, le noeud envoie le bloc compact immediatement sans attendre de requete. Les mineurs et les noeuds bien connectes utilisent generalement le mode haute bande passante pour minimiser la latence.

Le relais de blocs compacts reduit la bande passante de propagation de 90 % ou plus pour un bloc typique. Une propagation plus rapide donne aux petits mineurs une chance plus equitable, renforçant la decentralisation.
