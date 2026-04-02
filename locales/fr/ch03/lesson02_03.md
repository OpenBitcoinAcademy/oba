## Exigences en ressources

Executer un noeud complet necessite de l'espace disque, de la bande passante et du temps. Le telechargement initial de la blockchain depasse 500 Go en 2024 et augmente d'environ 150 Mo par jour.

Bitcoin Core peut fonctionner en mode elague, supprimant les anciennes donnees de blocs apres verification et conservant uniquement l'ensemble UTXO et les blocs recents. Cela reduit le stockage a moins de 10 Go. Le noeud verifie toujours l'historique complet lors de la synchronisation initiale.

Un noeud complet peut tourner sur du materiel aussi modeste qu'un Raspberry Pi. La synchronisation initiale prend plus de temps sur du materiel lent, mais une fois synchronise, les donnees quotidiennes sont gerable sur la plupart des connexions internet domestiques.
