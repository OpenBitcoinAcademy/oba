## Die Mining-Lotterie

Mining ist eine dezentrale Lotterie. Jeder Miner erstellt einen Block-Kandidaten aus Transactions in seinem Mempool. Er hasht den Block-Header wiederholt, ändert dabei jedes Mal eine Zahl namens Nonce und sucht nach einem Hash-Wert unterhalb eines vom Netzwerk festgelegten Zielwerts.

Einen gültigen Hash zu finden erfordert Billionen von Versuchen. Einen gültigen Hash zu überprüfen erfordert eine einzige Berechnung. Diese Asymmetrie ist der Kern von Proof of Work: Einen gültigen Block zu erzeugen ist teuer, ihn zu prüfen ist billig.

Der Miner, der zuerst einen gültigen Hash findet, verbreitet den Block. Andere Miner verifizieren ihn, akzeptieren ihn und beginnen sofort mit der Arbeit am nächsten Block. Der Gewinner erhält die Block-Belohnung: neu erzeugte Bitcoins (die Subvention) plus die Summe aller Transaktionsgebühren im Block.
