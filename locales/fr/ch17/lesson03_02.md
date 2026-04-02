## Trois couches : SimplPedPop, EncPedPop, ChillDKG

La fondation est SimplPedPop. Chaque participant genere son propre polynome aleatoire de degre t-1 et envoie une evaluation secrete a chaque autre participant, accompagnee d'un engagement sur les coefficients de son polynome. Chaque participant additionne les evaluations recues pour calculer sa part secrete finale. La cle publique du groupe est derivee de la somme des engagements de tous les participants sur leurs termes constants. Aucune partie ne detient la cle secrete complete.

SimplPedPop suppose un canal securise entre chaque paire de participants. EncPedPop ajoute cela en faisant generer a chaque participant une paire de cles de chiffrement ephemere et en chiffrant ses evaluations secretes avec la cle publique du destinataire. Le protocole fonctionne maintenant sur un canal de diffusion public, car les espions ne peuvent pas dechiffrer les parts secretes.

ChillDKG enveloppe EncPedPop avec une couche de gestion de session. Il introduit une cle secrete d'hote que chaque participant detient de maniere persistante, un jeu de donnees de recuperation commun identique pour tous les participants, et un protocole pour detecter et gerer les comportements malveillants. La cle secrete d'hote, combinee aux donnees de recuperation communes, permet a un participant de reconstruire sa part s'il perd son appareil de signature.
