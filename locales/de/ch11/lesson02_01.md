## Aufbau des Block-Headers

Jeder Block beginnt mit einem 80-Byte-Header, der sechs Felder enthält.

**Version** (4 Bytes): gibt an, welchen Konsensregeln der Block folgt, und signalisiert Unterstützung für Soft-Fork-Vorschläge über Versionbits.

**Previous Block Hash** (32 Bytes): der doppelte SHA-256-Hash des Headers des vorherigen Blocks. Das verknüpft jeden Block mit seinem Vorgänger und bildet die Chain.

**Merkle Root** (32 Bytes): ein einzelner Hash, der sich auf jede Transaction im Block festlegt. Das Ändern einer Transaction ändert die Merkle Root.

**Timestamp** (4 Bytes): der ungefähre Zeitpunkt, zu dem der Block gemined wurde, in Unix-Epochen-Sekunden.

**Target Bits** (4 Bytes): eine kompakte Kodierung des Proof-of-Work-Targets. Der Hash des Block-Headers muss unter diesem Target liegen.

**Nonce** (4 Bytes): der Wert, den Miner ändern, um einen gültigen Hash zu finden.
