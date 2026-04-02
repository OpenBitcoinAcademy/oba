## Transaction Fees

Bitcoin-Transactions haben kein explizites Gebührenfeld. Die Gebühr ist implizit: Sie ist die Differenz zwischen dem Gesamtwert aller Inputs und dem Gesamtwert aller Outputs.

Wenn deine Inputs insgesamt 100.000 Satoshis betragen und deine Outputs insgesamt 99.800 Satoshis, beträgt die Gebühr 200 Satoshis. Miner kassieren diese Gebühr als Anreiz, deine Transaction in einen Block aufzunehmen.

Höhere Gebühren bedeuten schnellere Bestätigung. Wenn das Netzwerk ausgelastet ist, priorisieren Miner Transactions mit höheren Gebührenraten. Eine Transaction, die zu wenig zahlt, wartet möglicherweise Stunden oder Tage.

Die Gebührenrate hängt vom Transaction-Gewicht ab (gemessen in virtuellen Bytes, oder vbytes), nicht vom gesendeten Betrag. Segwit führte Gewichtseinheiten ein, bei denen Witness-Daten (Signatures) gegenüber anderen Transaction-Daten vergünstigt werden. Eine Transaction, die 0,001 BTC sendet, kann dieselbe Gebühr kosten wie eine, die 1.000 BTC sendet, wenn beide dieselbe Struktur haben.
