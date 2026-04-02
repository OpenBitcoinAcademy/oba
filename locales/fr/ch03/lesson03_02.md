## Explorer la blockchain

Avec `bitcoin-cli`, vous pouvez recuperer n'importe quel bloc par sa hauteur ou son hachage. La commande `getblock` renvoie les champs de l'en-tete du bloc, la liste des identifiants de transaction et des metadonnees comme le nombre de confirmations.

La commande `getrawtransaction` renvoie les donnees brutes hexadecimales d'une transaction ou sa representation JSON decodee, montrant les entrees, les sorties, les scripts et les donnees temoin.

La commande `getblockchaininfo` rapporte la hauteur de chaine actuelle, la difficulte, la progression de la verification et si le noeud est toujours en cours de synchronisation. Ces commandes transforment les concepts abstraits de la blockchain en donnees concretes et inspectables.
