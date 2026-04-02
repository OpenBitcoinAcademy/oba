## Deterministische Key-Erzeugung

Deterministische Key-Erzeugung löst das Backup-Problem. Ein einzelner zufälliger Seed erzeugt eine Folge von Private Keys durch wiederholtes Hashing. Jeder Key in der Folge lässt sich aus dem Seed neu erstellen.

Sichere den Seed einmal, und du kannst jeden Key wiederherstellen, den das Wallet erzeugt hat oder erzeugen wird. Der Seed ist das Master-Geheimnis. Verlierst du ihn, verlierst du den Zugang zu allen abgeleiteten Keys.

Die einfachste Form: Starte mit einem Seed, hashe ihn zu Key 1, hashe Key 1 zu Key 2, und so weiter. Das ergibt eine flache Liste von Keys. Moderne Wallets verwenden einen ausgefeilteren Ansatz: einen Baum.
