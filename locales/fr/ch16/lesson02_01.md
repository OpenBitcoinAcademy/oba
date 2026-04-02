## Des roles, pas des personnes

Le flux de travail PSBT definit sept roles : Createur, Metteur a jour, Signataire, Combineur, Finaliseur, Extracteur et (dans PSBTv2) Constructeur. Chaque role effectue une etape. Une seule personne ou un seul programme peut remplir plusieurs roles, mais la separation compte car chaque role a des exigences de confiance differentes et un acces different aux informations.

Le Createur construit le PSBT initial. Dans v0 (BIP 174), le Createur produit la transaction non signee et l'enveloppe dans un PSBT avec des cartes d'entrees et de sorties vides. Dans v2 (BIP 370), le Createur definit les champs globaux comme la version de transaction et le locktime, mais n'inclut pas encore les entrees ou sorties. Cette tache revient au Constructeur.

Le role de Constructeur n'existe que dans PSBTv2. Il ajoute des entrees et des sorties au PSBT. Plusieurs Constructeurs peuvent collaborer : l'un ajoute les entrees qu'il controle, un autre ajoute les siennes, et chacun ajoute les sorties dont il a besoin. Cela permet la construction interactive de transactions ou aucune partie ne connait la forme complete de la transaction tant que tous les Constructeurs n'ont pas contribue.
