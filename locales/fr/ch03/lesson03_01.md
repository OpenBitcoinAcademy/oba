## L'interface JSON-RPC

Bitcoin Core expose une interface JSON-RPC qui permet aux programmes d'interroger les donnees de la blockchain et de soumettre des transactions. L'outil en ligne de commande `bitcoin-cli` envoie des requetes a cette interface.

Chaque donnee de la blockchain est interrogeable : en-tetes de blocs, blocs complets, transactions individuelles, soldes d'adresses, contenu du mempool, connexions aux pairs et statistiques du reseau.

Les portefeuilles, explorateurs de blocs, processeurs de paiement et outils de recherche communiquent tous avec Bitcoin Core via cette API. L'interface est le pont entre les donnees brutes de la blockchain et les applications qui les presentent aux utilisateurs.
