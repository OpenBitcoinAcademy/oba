## SegWits Signing-Algorithmus

Segwit führte einen neuen Signing-Algorithmus ein (BIP143), der ein langjähriges Problem behebt: Bei Legacy-Transactions ist der ausgegebene Betrag nicht in den signierten Daten enthalten. Das zwang Unterzeichner, die vollständige vorherige Transaction abzurufen, um den Betrag zu verifizieren, was Hardware Wallets verlangsamte.

BIP143 nimmt den Betrag jedes Inputs in die signierten Daten auf. Ein Hardware Wallet kann den gesamten Input-Wert und die Gebühr verifizieren, ohne vorherige Transactions herunterzuladen. Der Signing-Prozess ist schneller und sicherer.

Für Segwit v1 (Taproot) definiert BIP341 einen aktualisierten Algorithmus mit zusätzlichen Commitments, der sowohl Key Path als auch Script Path Spending unterstützt.
