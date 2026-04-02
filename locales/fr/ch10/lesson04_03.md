## Filtres de blocs compacts

**BIP 157** et **BIP 158** definissent une meilleure approche appelee Compact Block Filters. Au lieu que le client envoie son filtre au serveur, le serveur construit un filtre pour chaque bloc et l'envoie au client.

Chaque filtre est une representation compacte de toutes les adresses et scripts dans un bloc. Le client telecharge le filtre et verifie localement si l'une de ses adresses apparait. Si une correspondance est trouvee, le client telecharge le bloc complet pour extraire les transactions pertinentes.

L'avantage en matiere de vie privee est que le serveur n'apprend jamais quelles adresses interessent le client. Le serveur envoie le meme filtre a chaque client. Le client fait toute la correspondance localement.

Le cout en bande passante est plus eleve que les filtres de Bloom car le client telecharge un filtre pour chaque bloc. Mais les filtres sont petits (environ 20 Ko par bloc en moyenne) et peuvent etre verifies par rapport a un engagement de hachage dans la chaine d'en-tetes de blocs. Le client n'a pas besoin de faire confiance au serveur pour fournir des filtres precis.

Faire tourner Bitcoin sur **Tor** ajoute une couche supplementaire de vie privee. Tor cache l'adresse IP du client aux noeuds auxquels il se connecte. Combine avec les Compact Block Filters, un client leger peut verifier son solde sans reveler son identite ou ses adresses a aucun pair.
