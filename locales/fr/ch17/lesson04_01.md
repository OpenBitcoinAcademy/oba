## La reutilisation de nonce est catastrophique

FROST herite de la sensibilite de Schnorr a la gestion des nonces. Si un signataire utilise le meme nonce pour deux sessions de signature differentes, un attaquant qui observe les deux parts de signature peut extraire la part de cle secrete de ce signataire. Avec suffisamment de parts extraites (t d'entre elles), l'attaquant peut reconstruire la cle secrete complete du groupe et voler tous les fonds.

Ce n'est pas une preoccupation theorique. La generation deterministe de nonces, qui est sure pour Schnorr a signataire unique, est dangereuse dans le contexte a seuil. Si un coordinateur malveillant rejoue une ancienne requete de signature, un signataire utilisant des nonces deterministes produirait une nouvelle part de signature avec le meme nonce, divulguant sa part. FROST exige donc des nonces cryptographiquement aleatoires pour chaque session de signature.

Les implementations doivent generer les nonces a partir d'une source aleatoire forte et ne jamais persister les nonces entre les sessions. Si une session de signature est abandonnee, les nonces utilises dans cette session doivent etre supprimes. Re-entrer dans une session avec les memes nonces equivaut a une reutilisation de nonce.
