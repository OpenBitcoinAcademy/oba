## Transmission multisig et transport

Dans une configuration multisig, le PSBT passe par chaque signataire en sequence ou est distribue a tous les signataires en parallele. Considerez un multisig 2-of-3 : le coordinateur cree le PSBT et l'envoie au signataire A et au signataire B simultanement. Chaque signataire ajoute sa signature partielle et renvoie le PSBT. Le coordinateur combine les deux PSBT, finalise et diffuse.

Le mecanisme de transport n'est pas specifie par BIP 174. Les PSBT peuvent voyager sur des cartes SD, comme des codes QR (en utilisant l'encodage UR de BCR-2020-005), par NFC, par courriel, via un service de partage de fichiers ou par tout autre canal. Le format est autonome. Chaque signataire obtient tout ce dont il a besoin du PSBT lui-meme, sans canal lateral requis.

PSBT est en clair, pas chiffre. Quiconque intercepte un PSBT peut voir les montants de la transaction, les adresses impliquees et les signatures partielles collectees. C'est un probleme de vie privee, pas un probleme de securite. Un attaquant qui lit un PSBT ne peut pas voler de fonds car le PSBT ne contient pas de cles privees. Mais l'attaquant apprend ce que l'utilisateur depense et ou vont les fonds. Pour les transactions sensibles, le PSBT devrait etre transporte sur un canal chiffre ou sur un support physique que l'utilisateur controle.
