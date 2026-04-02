## Abandons identifiables et MuSig2 contre FROST

Un participant malveillant peut perturber la signature en soumettant une part de signature invalide. Sans verifications supplementaires, le coordinateur additionnerait les parts et produirait une signature finale invalide, mais ne saurait pas quel signataire s'est mal comporte. FROST supporte les abandons identifiables : le coordinateur peut verifier chaque part de signature individuellement par rapport a la part de cle publique du signataire. Le signataire malveillant est identifie et peut etre exclu des sessions futures.

MuSig2 et FROST repondent a des besoins differents. MuSig2 est un schema n-of-n : tous les participants doivent signer, il n'y a pas de seuil, et le protocole est plus elegante. FROST est un schema t-of-n : il tolere les signataires absents mais necessite une phase de generation de cle plus complexe. Les deux produisent une sortie sur la chaine identique : une seule signature Schnorr de 64 octets sous une cle publique de 32 octets.

MuSig2 convient aux scenarios ou tous les signataires sont censes etre disponibles, comme un canal Lightning entre deux noeuds. FROST convient aux arrangements de garde ou la redondance compte, comme une tresorerie d'entreprise avec cinq detenteurs de cles ou trois quelconques peuvent autoriser une depense.
