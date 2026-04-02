## Limites supprimees

Le Script legacy et SegWit v0 imposent des limites de ressources strictes : taille de script de 10 000 octets, 201 opcodes non-push, 100 elements de pile temoin. Ces limites etaient necessaires car le modele de cout ne prenait pas en compte les scripts individuels.

Tapscript supprime la limite de taille de script et la limite de nombre d'opcodes. A la place, il utilise un budget d'operations de signature par entree : 50 sigops + le poids temoin de l'entree. Les temoins plus gros paient des frais plus eleves, et le budget de sigops s'adapte aux frais payes. L'incitation economique remplace la limite fixe.

Les scripts qui etaient impossibles sous les anciennes limites fonctionnent maintenant. Un seuil avec 100 participants, une cascade complexe avec verrous temporels ou une verification de chaine de hachage peuvent etre exprimes dans une seule feuille Tapscript.
