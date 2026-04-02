## Generation aleatoire de cles

Les premiers portefeuilles Bitcoin generaient chaque cle privee independamment a l'aide d'un generateur de nombres aleatoires. Chaque cle etait sans lien avec les autres. Le portefeuille les stockait toutes dans un fichier de base de donnees.

Cette approche posait un serieux probleme de sauvegarde. Chaque nouvelle cle necessitait une nouvelle sauvegarde. Si un utilisateur generait 100 cles et faisait une sauvegarde apres la cle 50, les cles 51 a 100 seraient perdues en cas de corruption du fichier de portefeuille.

Bitcoin Core generait a l'origine un pool de 100 cles a l'avance pour reduire la frequence des sauvegardes requises. Mais le probleme persistait : les cles aleatoires independantes sont difficiles a gerer.
