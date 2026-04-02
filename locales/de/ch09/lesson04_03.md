## Fee Sniping

Fee Sniping ist ein theoretischer Angriff, bei dem ein Miner, anstatt auf dem neuesten Block aufzubauen, einen vorherigen Block erneut mined, um dessen Transaction Fees zu stehlen. Wenn die Fees in Block N hoch genug sind, könnte ein Miner es profitabler finden, die Chain bei Block N-1 zu forken und diese Fees selbst zu kassieren.

Bitcoin Core verteidigt dagegen, indem es die Locktime neuer Transactions auf die aktuelle Block-Höhe setzt. Eine Transaction, die auf Block 800.000 gesperrt ist, kann nicht in Block 799.999 aufgenommen werden. Das macht das erneute Mining alter Blocks weniger profitabel, weil die seitdem erstellten neuen Transactions nicht verfügbar wären.

Fee Sniping hat im Bitcoin-Netzwerk nicht stattgefunden. Die Verteidigung ist vorbeugend.
