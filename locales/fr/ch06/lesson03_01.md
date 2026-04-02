## Le format filaire

Une transaction Bitcoin est serialisee en une sequence d'octets pour la transmission a travers le reseau et le stockage dans les blocs. Le format de serialisation definit l'ordre exact des octets et l'encodage de chaque champ.

Une transaction legacy a quatre champs de haut niveau serialises dans l'ordre : version (4 octets, little-endian), entrees (variable), sorties (variable) et locktime (4 octets, little-endian).

Une transaction segwit ajoute trois champs : apres la version, un octet marqueur (0x00) et un octet drapeau (0x01) signalent que des donnees temoin suivent. La structure temoin apparait apres les sorties et avant le locktime. Un analyseur qui voit un octet zero la ou le compteur d'entrees devrait etre sait qu'il lit le format etendu (segwit).
