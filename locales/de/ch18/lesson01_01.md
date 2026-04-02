## Zwei Transactions für Tausende Zahlungen

Ein Payment Channel ist ein 2-of-2-Multisig-Vertrag auf der Bitcoin-Blockchain. Zwei Parteien sperren Mittel in einen gemeinsamen Output. Sie tauschen signierte Commitment Transactions off-chain aus, um die Aufteilung des Guthabens zu aktualisieren. Nur zwei On-Chain-Transactions sind nötig: eine zum Öffnen des Channels und eine zum Schließen.

Zwischen Öffnen und Schließen können die Parteien Tausende Zahlungen austauschen. Jede Zahlung aktualisiert die Commitment Transaction, die die gesperrten Mittel aufteilt. Die Blockchain sieht diese Zwischenstände nie. Das Ergebnis: schnelle Zahlungen (Millisekunden), nahezu null Gebühren und keine Wartezeit auf Block-Bestätigungen.
