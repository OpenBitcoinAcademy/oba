## Mining als dezentrale Lotterie

Bitcoin hat keine zentrale Instanz, die entscheidet, welche Transactions in Blocks kommen. Stattdessen konkurrieren Miner in einer rechnerischen Lotterie. Der Gewinner erhält das Recht, den nächsten Block an die Chain anzuhängen.

Jeder Miner nimmt eine Menge unbestätigter Transactions, stellt sie zu einem Kandidatenblock zusammen und hasht den Block-Header mit SHA-256 (doppelt angewendet). Das Ergebnis muss eine Zahl unter dem aktuellen Target sein. Die meisten Versuche scheitern. Der Miner ändert das Nonce-Feld im Header und hasht erneut.

Dieser Vorgang wiederholt sich milliardenfach pro Sekunde über alle Miner weltweit. Der Miner, der zuerst einen gültigen Hash findet, sendet den Block. Andere Nodes verifizieren ihn unabhängig. Niemand braucht eine Erlaubnis zum Mining, und niemand kann vorhersagen, wer als nächstes gewinnt.
