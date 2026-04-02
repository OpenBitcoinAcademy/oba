## Die Coinbase-Transaction

Jeder Block beginnt mit einer speziellen Transaction, der Coinbase-Transaction. Sie hat keine Inputs aus vorherigen Transactions. Stattdessen erzeugt sie neue Bitcoin als Belohnung für den Miner, der den Block gefunden hat.

Die Coinbase-Transaction hat einen Input mit einer Null-Referenz (kein vorheriger Output wird ausgegeben). Der Miner füllt das Script-Feld dieses Inputs mit beliebigen Daten, einschließlich der Extra Nonce, die beim Mining verwendet wird. Satoshi Nakamoto hat eine Zeitungsschlagzeile in die Coinbase des Genesis Blocks eingebettet: "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."

Die Outputs der Coinbase-Transaction zahlen den Miner. Der Gesamtwert der Outputs darf die Block-Subsidy plus die Summe aller Transaction Fees im Block nicht übersteigen. Nicht beanspruchte Werte gehen verloren.
