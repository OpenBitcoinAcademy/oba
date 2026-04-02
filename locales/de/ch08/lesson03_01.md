## Schnorr: Eine einfachere Signature

Schnorr-Signatures (BIP340) wurden mit dem Taproot-Upgrade im November 2021 aktiviert. Sie werden für Segwit v1 (P2TR) Transactions verwendet.

Der Schnorr-Algorithmus ist mathematisch einfacher als ECDSA. Er erzeugt eine feste 64-Byte-Signature (im Vergleich zu ECDSAs variablen 71-73 Bytes). Die Verifikationsgleichung ist linear, was Signature-Aggregation ermöglicht: Mehrere Signatures können zu einer einzelnen Signature derselben Größe kombiniert werden.

Schnorr-Signatures haben eine beweisbare Sicherheitsreduktion auf das Diskrete-Logarithmus-Problem, eine stärkere Sicherheitsgarantie als ECDSA.
