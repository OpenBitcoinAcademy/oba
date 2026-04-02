## Zeitverfallende Ausgabepolicies

Eine Policy mit zeitgesperrter Wiederherstellung beginnt mit einer primären Ausgabebedingung und fügt Rückfallbedingungen hinzu, die nach einer Verzögerung aktiviert werden. Der primäre Key kontrolliert die Mittel sofort. Geht der primäre Key verloren oder ist sein Inhaber nicht erreichbar, wird ein alternativer Ausgabepfad nach einer bestimmten Anzahl von Blocks freigeschaltet.

In Miniscript wird das ausgedrückt als: `or(pk(primary),and(pk(recovery),older(26280)))`. Der primäre Key kann jederzeit ausgeben. Der Recovery-Key kann erst nach ungefähr sechs Monaten ausgeben (26.280 Blocks bei 10 Minuten pro Block). Die Policy kompiliert zu einem einzigen Script mit zwei Ausführungspfaden.

Dieses Muster löst ein reales Problem. Ein Einzel-Key-Wallet hat keinen Wiederherstellungspfad. Geht der Key verloren, sind die Mittel weg. Ein Standard-Multisig erfordert mehrere Signer für jede Transaction, was den täglichen Gebrauch umständlich macht. Ein zeitgesperrter Rückfall gibt dem primären Inhaber volle Kontrolle im Normalbetrieb, während er garantiert, dass eine vertrauenswürdige Partei die Mittel nach einer Verzögerung wiederherstellen kann.
