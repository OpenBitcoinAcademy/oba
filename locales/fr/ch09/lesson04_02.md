## Transaction Pinning

Le transaction pinning est une attaque ou une partie malveillante cree une transaction enfant a frais bas qui rend couteux ou impossible pour la partie honnete d'augmenter les frais du parent.

Dans un protocole a deux parties (comme un canal Lightning Network), une partie pourrait diffuser un descendant volumineux et a frais bas d'une transaction partagee. La tentative de CPFP de l'autre partie devrait payer pour le descendant volumineux de l'attaquant, rendant l'augmentation de frais prohibitivement couteuse.

Les sorties d'ancrage sont une contre-mesure. Chaque partie obtient une petite sortie dans la transaction partagee qu'elle peut depenser avec CPFP. Des regles limitent le nombre de descendants que chaque ancrage peut avoir, empechant l'attaque par pinning.
