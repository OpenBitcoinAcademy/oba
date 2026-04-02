## Die Transaction erstellen

Eine Legacy-Transaction hat vier Felder: Version, Inputs, Outputs und Locktime. Moderne Segwit-Transactions fügen drei weitere hinzu: einen Marker, ein Flag und eine Witness-Struktur, die die Autorisierungsdaten (Signatures) getrennt von den Inputs hält.

**Version** ist eine Zahl (derzeit 1 oder 2), die den Nodes mitteilt, welche Validierungsregeln gelten.

**Inputs** listen die auszugebenden UTXOs auf. Jeder Input gibt die vorherige Transaction-ID, den Output-Index in dieser Transaction, ein Input Script und eine Sequence Number an.

**Outputs** listen die neu erzeugten UTXOs auf. Jeder Output gibt einen Betrag in Satoshis und ein Locking Script an.

**Locktime** ist normalerweise null. Wenn auf eine zukünftige Blockhöhe oder einen Zeitstempel gesetzt, kann die Transaction erst ab diesem Punkt in einen Block aufgenommen werden.

Die Transaction wird in eine Byte-Folge serialisiert, zweimal mit SHA-256 gehasht, und das Ergebnis ist die Transaction-ID.
