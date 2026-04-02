## Schnorr vs ECDSA

Beide Algorithmen verwenden die secp256k1-Kurve. Beide bieten gleichwertige Sicherheit gegen Key-Wiederherstellung. Die Unterschiede sind praktischer Natur.

Schnorr-Signatures sind kleiner (64 Bytes gegenüber 71-73). Schnorr unterstützt native Signature-Aggregation. Schnorr hat eine einfachere Verifikationsgleichung. Schnorr hat einen formalen Sicherheitsbeweis.

ECDSA wurde 2008 für Bitcoin gewählt, weil Schnorr bis 2008 patentiert war und der Patentstatus unklar war, als Satoshi das System entwarf. Nach Ablauf des Patents und der Aktivierung von Taproot können neue Transactions Schnorr verwenden. ECDSA wird weiterhin aus Gründen der Rückwärtskompatibilität unterstützt.
