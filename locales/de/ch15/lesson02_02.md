## Das Typsystem

Jedes Miniscript-Fragment hat einen Typ, der sein Verhalten auf dem Script-Stack beschreibt. Die vier Basistypen sind B, V, K und W.

Ein B-Typ (Base) Ausdruck legt bei Erfolg einen einzelnen Nicht-Null-Wert auf den Stack und Null bei Fehlschlag. Ein V-Typ (Verify) Ausdruck ist entweder erfolgreich ohne Stack-Ausgabe oder bricht das Script ab. Ein K-Typ (Key) Ausdruck legt bei Erfolg einen Public Key auf den Stack. Ein W-Typ (Wrapped) Ausdruck verhält sich wie B, konsumiert aber zuerst das oberste Stack-Element; er wird zum Kombinieren von Teilausdrücken verwendet.

Der Compiler prüft Typen an jeder Kompositionsstelle. Ein `and_v`-Fragment verlangt, dass sein erstes Argument V-Typ ist und sein zweites B-Typ. Übergibst du Argumente mit falschen Typen, lehnt der Compiler den Ausdruck ab. Das fängt Fehler ab, die in rohem Script ein Script erzeugen würden, das "funktioniert", aber die falschen Bedingungen durchsetzt.

Das Typsystem verfolgt auch Eigenschaften wie Dissipierbarkeit (kann ein fehlgeschlagener Zweig den Stack sauber hinterlassen?) und Nicht-Manipulierbarkeit (gibt es für jeden Ausgabepfad genau eine gültige Witness?). Diese Eigenschaften pflanzen sich durch die Komposition fort. Ein Miniscript-Ausdruck hat sie oder nicht, und der Compiler kann berichten, welche vorliegen.
