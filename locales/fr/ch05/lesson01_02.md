## La base de donnees du portefeuille

Une application de portefeuille stocke ses donnees dans une base de donnees de portefeuille. Cette base de donnees contient les cles privees, les cles publiques et des metadonnees comme les libelles de transactions, les notes d'adresses et les parametres de frais.

Certains portefeuilles ne stockent que les cles publiques et dependent d'appareils externes (portefeuilles materiels) pour detenir les cles privees et produire les signatures. D'autres stockent toutes les cles localement. Le modele de securite differe, mais le principe est le meme : la base de donnees du portefeuille est le lien entre l'utilisateur et la blockchain.

Les applications de portefeuille fournissent une interface utilisateur par-dessus la base de donnees : affichage des soldes, construction des transactions, gestion des adresses et connexion au reseau Bitcoin.
