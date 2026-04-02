## Vom Mempool zur Blockchain

Alice' Transaction befindet sich im Mempool und wartet darauf, dass ein Miner sie in einen Block aufnimmt. Miner wählen Transactions aus ihrem Mempool, wobei sie solche mit höheren Gebührenraten pro Gewichtseinheit bevorzugen.

Etwa fünf Minuten später findet ein Miner eine gültige Proof-of-Work-Lösung für einen neuen Block, der Alice' Transaction enthält. Der Miner verbreitet den Block. Jede Full Node verifiziert den Block unabhängig: gültiger Proof of Work, gültige Transactions, korrektes Format. Nodes, die den Block akzeptieren, fügen ihn ihrer Kopie der Blockchain hinzu.

Alice' Transaction hat jetzt eine Bestätigung. Bobs Wallet aktualisiert den Status. Jeder folgende Block, der auf diesem aufbaut, fügt eine weitere Bestätigung hinzu. Nach sechs Bestätigungen (etwa eine Stunde) gilt die Transaction als hochgradig sicher gegen Umkehrung.
