## Die Multisig-Landschaft

Bitcoin unterstützt seit 2012 die Verwahrung durch mehrere Parteien per OP_CHECKMULTISIG. Ein 2-of-3-Multisig-Script listet alle drei Public Keys on-chain auf und verlangt zwei gültige Signatures. Das funktioniert, legt aber offen, wie viele Parteien die Mittel kontrollieren. Jeder, der die Blockchain inspiziert, sieht den Schwellenwert, die Gesamtzahl der Unterzeichner und jeden beteiligten Public Key.

Die On-Chain-Kosten skalieren linear. Ein 3-of-5-Multisig platziert fünf 33-Byte-Public-Keys und drei 72-Byte-Signatures in den Witness. Das sind 381 Bytes Witness-Daten, bevor das Script selbst mitgezählt wird. Ein 7-of-10-Setup kostet noch mehr. Gebühren steigen, Privatsphäre sinkt, und die Spending Policy ist für alle Beobachter sichtbar.

Taproot und Schnorr-Signatures haben die Möglichkeiten verändert. MuSig2 erlaubt n Teilnehmern, eine einzige aggregierte Signature zu erzeugen, die on-chain identisch zu einer Solo-Schnorr-Signature aussieht. Der kombinierte Public Key ist 32 Bytes groß. Die Signature ist 64 Bytes groß. Kein Beobachter kann erkennen, wie viele Parteien beteiligt waren.
