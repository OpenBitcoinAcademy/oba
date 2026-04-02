## Signatures in Bitcoin-Transactions

In Bitcoin erscheinen Signatures in Transaction-Inputs (bei Legacy-Transactions) oder in der Witness-Struktur (bei Segwit-Transactions). Jeder Input, der einen UTXO ausgibt, braucht eine Signature, die beweist, dass der Absender den Key kontrolliert, der diesen Output gesperrt hat.

Mehrere Parteien können an einer einzelnen Transaction zusammenarbeiten. Jede Partei signiert nur den Input, den sie kontrolliert. Die Signatures sind unabhängig: Ein Unterzeichner braucht keinen Zugang zum Private Key eines anderen.

Bitcoin verwendet zwei Signature-Algorithmen: ECDSA (Elliptic Curve Digital Signature Algorithm) für Legacy- und Segwit-v0-Transactions, und Schnorr-Signatures für Segwit v1 (Taproot) Transactions.
