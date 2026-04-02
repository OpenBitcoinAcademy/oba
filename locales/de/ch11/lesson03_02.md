## Merkle Proofs

Ein Merkle Proof ermöglicht die Verifizierung, dass eine bestimmte Transaction in einem Block enthalten ist, ohne jede Transaction herunterzuladen. Der Proof besteht aus dem Transaction-Hash, den Geschwister-Hashes entlang des Pfads zur Wurzel und der Merkle Root aus dem Block-Header.

Zum Verifizieren: Beginne mit dem Transaction-Hash. Hashe ihn mit seinem Geschwister. Hashe das Ergebnis mit dem nächsten Geschwister. Setze fort, bis du die Wurzel erreichst. Wenn deine berechnete Wurzel mit der Merkle Root im Block-Header übereinstimmt, ist die Transaction im Block enthalten.

Ein Block mit 4.000 Transactions erfordert einen Merkle Proof von nur etwa 12 Hashes (log2(4000)). SPV-Clients nutzen Merkle Proofs, um Transactions zu verifizieren, ohne vollständige Blocks herunterzuladen.
