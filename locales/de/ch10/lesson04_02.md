## Bloom Filter und ihr Datenschutzproblem

**BIP 37** führte Bloom Filter ein, damit SPV-Clients nur relevante Transactions von Full Nodes anfordern können. Ein Bloom Filter ist eine probabilistische Datenstruktur. Der Client erstellt einen Filter mit seinen Addresses und sendet ihn an einen Peer. Der Peer sendet nur Transactions, die zum Filter passen.

Das Design hat einen schwerwiegenden Datenschutzmangel. Die Full Node sieht den Bloom Filter und kann ableiten, welche Addresses zum Client gehören. Obwohl ein Bloom Filter Falsch-Positive hat, zeigt die Forschung, dass eine böswillige Node die Addresses des Clients mit hoher Genauigkeit identifizieren kann, indem sie das Bitmuster des Filters analysiert.

Eine Node, die Bloom-Filter-Anfragen bedient, trägt auch hohe Rechenkosten. Sie muss jede Transaction in jedem Block gegen den Filter testen. Das erzeugt einen Denial-of-Service-Vektor: Ein Angreifer kann sich vielfach mit verschiedenen Filtern verbinden und die Node zu teurer Arbeit zwingen. Viele Node-Betreiber deaktivieren BIP-37-Unterstützung.
