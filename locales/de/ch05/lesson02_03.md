## Vom Seed zum Baum

Ein Seed ist eine Zufallszahl, typischerweise 128 bis 256 Bit. Er ist die Wurzel aller Key-Ableitung. Aus diesem Seed kann eine Hierarchie von Keys erzeugt werden: Der Seed erzeugt einen Master Key, der Master Key erzeugt Child Keys, jeder Child kann eigene Children erzeugen.

Diese Baumstruktur ist die Basis von Hierarchical Deterministic (HD) Wallets, definiert in BIP32. Der Baum erlaubt die Organisation von Keys nach Zweck: ein Zweig für den Empfang von Zahlungen, ein anderer für Wechselgeld, ein weiterer für ein anderes Konto.

Der Seed ändert sich nie. Jeder Key im Baum lässt sich daraus neu erzeugen. Ein einzelnes Backup schützt eine unbegrenzte Anzahl zukünftiger Keys.
