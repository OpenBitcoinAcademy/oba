## Le repertoire de donnees

Bitcoin Core stocke ses donnees dans un repertoire specifique a la plateforme : `~/.bitcoin` sur Linux, `~/Library/Application Support/Bitcoin` sur macOS, `%APPDATA%\Bitcoin` sur Windows.

Le sous-repertoire `blocks/` contient les fichiers de donnees brutes des blocs. Le repertoire `chainstate/` contient la base de donnees UTXO, un stockage LevelDB de toutes les sorties non depensees. Le fichier `mempool.dat` persiste le mempool entre les redemarrages.

Le fichier de configuration `bitcoin.conf` controle les parametres reseau, les limites de ressources, l'authentification RPC et les options de fonctionnalites. Des parametres comme `prune=550` activent le mode elague, et `txindex=1` construit un index de toutes les transactions pour des recherches plus rapides.
