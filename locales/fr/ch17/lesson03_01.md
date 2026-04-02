## Le probleme du distributeur de confiance

La facon la plus elegante de configurer les parts de cle FROST est avec un distributeur de confiance. Une partie genere la cle secrete du groupe, evalue le polynome de Shamir en n points et distribue une part a chaque participant. Cela fonctionne mais cree un point de defaillance unique : le distributeur connait la cle secrete complete. Si le distributeur est compromis, les fonds du groupe entier sont en danger.

Un distributeur compromis peut aussi distribuer des parts inconsistantes, donnant a certains participants des donnees invalides qui feront echouer la signature de maniere imprevisible. Les participants n'ont aucun moyen de verifier que leurs parts sont correctes sans un protocole qui impose la coherence.

ChillDKG, specifie dans BIP 445, elimine le distributeur de confiance. C'est un protocole de generation de cle distribuee : le groupe produit conjointement les parts de cle sans qu'aucune partie ne detienne ou voie la cle secrete complete. Le protocole se construit en trois couches, chacune ajoutant une garantie specifique.
