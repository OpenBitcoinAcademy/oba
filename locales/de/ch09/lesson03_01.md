## Replace By Fee (RBF)

Wenn eine Transaction im Mempool feststeckt, weil ihre Fee-Rate zu niedrig ist, kann der Absender sie durch eine neue Version mit höherer Fee ersetzen. Das ist Replace By Fee (RBF), definiert in BIP125.

Für RBF muss die ursprüngliche Transaction Ersetzbarkeit signalisieren, indem die Sequence-Nummer mindestens eines Inputs auf einen Wert unter 0xFFFFFFFE gesetzt wird. Wallet-Software erledigt das automatisch.

Die Ersatz-Transaction muss eine höhere Gesamt-Fee zahlen als das Original. Sie kann die Outputs ändern (einen anderen Betrag zahlen oder Empfänger hinzufügen/entfernen), solange sie mindestens einen derselben Inputs ausgibt.
