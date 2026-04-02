## Taproot (P2TR)

Taproot (BIP341) ist der neueste Output-Typ, aktiviert im November 2021 als Segwit Version 1. Ein Taproot-Output bindet sich an einen einzelnen Public Key. Das Ausgeben kann auf zwei Wegen erfolgen.

**Key Path**: Der Besitzer signiert direkt mit dem gebundenen Key. Auf der Blockchain sieht das identisch aus wie eine Single-Signature-Ausgabe. Kein Script wird enthüllt. Niemand kann erkennen, ob der Output zusätzliche Ausgabebedingungen hatte.

**Script Path**: Der Besitzer enthüllt ein Script aus einem Merkle Tree von Scripts, die im Output gebunden sind. Das ermöglicht komplexe Bedingungen (Multisig, Timelocks, Hash Locks), während der häufige Fall (Key Path) kompakt und privat bleibt.

Taproot-Addresses beginnen mit "bc1p" im Mainnet und verwenden Bech32m-Kodierung.
