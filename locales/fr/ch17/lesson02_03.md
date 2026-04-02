## Des parts a une signature standard

Le coordinateur collecte les t parts de signature et les additionne. Le resultat est une seule signature Schnorr : une paire (R, s) de 64 octets, ou R est le point de nonce agrege et s est la somme des parts de signature. Cette signature est valide sous la cle publique agregee du groupe.

Un noeud complet Bitcoin qui verifie cette transaction execute l'equation de verification Schnorr BIP 340 standard. Il verifie une signature contre une cle publique. Le noeud n'a aucun moyen de savoir si la signature a ete produite par trois signataires, cinq signataires ou un seul signataire. La verification est identique.

C'est pourquoi FROST est puissant pour Bitcoin : il ne necessite aucun changement de consensus. Taproot et BIP 340 acceptent deja les signatures que FROST produit. La complexite de la signature a seuil reside entierement dans le logiciel des signataires. La blockchain et tous les verificateurs restent inconscients.
