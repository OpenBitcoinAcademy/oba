## Multisignature-Aggregation

Bei ECDSA braucht eine Transaction, die drei Unterzeichner erfordert, drei separate Signatures von je 71-73 Bytes. Die Signatures werden einzeln verifiziert.

Bei Schnorr können dieselben drei Unterzeichner ihre Signatures mit Protokollen wie MuSig2 zu einer einzigen 64-Byte-Aggregat-Signature kombinieren. Auf der Blockchain sieht das Aggregat identisch aus wie eine Single-Signer-Signature. Kein Beobachter kann erkennen, wie viele Parteien beteiligt waren.

Das hat zwei Vorteile. Transactions mit mehreren Unterzeichnern verbrauchen weniger Block Space (niedrigere Gebühren). Und Multisignature-Transactions werden von Single-Signature-Transactions ununterscheidbar, was die Privatsphäre verbessert.
