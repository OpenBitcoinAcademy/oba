## Construction de la transaction

Une transaction legacy a quatre champs : version, entrees, sorties et locktime. Les transactions segwit modernes ajoutent trois champs supplementaires : un marqueur, un drapeau et une structure temoin qui contient les donnees d'autorisation (signatures) separement des entrees.

**Version** est un nombre (actuellement 1 ou 2) qui indique aux noeuds quelles regles de validation appliquer.

**Entrees** listent les UTXO depenses. Chaque entree specifie l'identifiant de la transaction precedente, l'index de sortie dans cette transaction, un script d'entree et un numero de sequence.

**Sorties** listent les nouveaux UTXO crees. Chaque sortie specifie un montant en satoshis et un script de verrouillage.

**Locktime** est generalement zero. Quand il est regle sur une hauteur de bloc ou un horodatage futur, la transaction ne peut pas etre incluse dans un bloc avant ce point.

La transaction est serialisee en une sequence d'octets, hachee deux fois avec SHA-256, et le resultat est l'identifiant de transaction.
