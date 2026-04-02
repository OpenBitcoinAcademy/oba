## Identifizierbare Abbrüche und MuSig2 vs FROST

Ein bösartiger Teilnehmer kann das Signieren stören, indem er einen ungültigen Signature Share einreicht. Ohne zusätzliche Prüfungen würde der Coordinator die Shares summieren und eine ungültige finale Signature erzeugen, aber nicht wissen, welcher Unterzeichner sich fehlverhalten hat. FROST unterstützt identifizierbare Abbrüche: Der Coordinator kann jeden Signature Share einzeln gegen den Public Key Share des Unterzeichners verifizieren. Der Fehlverhalten zeigende Unterzeichner wird identifiziert und kann von künftigen Sitzungen ausgeschlossen werden.

MuSig2 und FROST bedienen unterschiedliche Bedürfnisse. MuSig2 ist ein n-of-n-Schema: Alle Teilnehmer müssen signieren, es gibt keinen Threshold, und das Protokoll ist einfacher. FROST ist ein t-of-n-Schema: Es toleriert abwesende Unterzeichner, erfordert aber eine komplexere Key-Generierungsphase. Beide erzeugen identischen On-Chain-Output: eine einzelne 64-Byte-Schnorr-Signature unter einem 32-Byte-Public-Key.

MuSig2 eignet sich für Szenarien, in denen alle Unterzeichner erreichbar sein dürften, etwa ein Lightning Channel zwischen zwei Nodes. FROST eignet sich für Verwahrungsarrangements, bei denen Redundanz zählt, etwa ein Unternehmens-Treasury mit fünf Schlüsselinhabern, von denen beliebige drei einen Spend autorisieren können.
