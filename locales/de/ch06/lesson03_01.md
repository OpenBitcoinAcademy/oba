## Das Wire Format

Eine Bitcoin-Transaction wird in eine Byte-Folge serialisiert, um über das Netzwerk übertragen und in Blocks gespeichert zu werden. Das Serialisierungsformat definiert die exakte Byte-Reihenfolge und Kodierung jedes Feldes.

Eine Legacy-Transaction hat vier Top-Level-Felder, die der Reihe nach serialisiert werden: Version (4 Bytes, Little-Endian), Inputs (variabel), Outputs (variabel) und Locktime (4 Bytes, Little-Endian).

Eine Segwit-Transaction fügt drei Felder hinzu: Nach der Version signalisieren ein Marker-Byte (0x00) und ein Flag-Byte (0x01), dass Witness-Daten folgen. Die Witness-Struktur erscheint nach den Outputs und vor der Locktime. Ein Parser, der ein Null-Byte sieht, wo die Input-Anzahl stehen sollte, erkennt, dass er das erweiterte (Segwit-)Format liest.
