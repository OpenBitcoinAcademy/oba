## Privatsphäre durch Ununterscheidbarkeit

Der Privatsphäregewinn durch Threshold Signatures geht über das Verbergen der Spending Policy hinaus. Bei OP_CHECKMULTISIG ist das Script-Template selbst ein Fingerabdruck. Chain-Analyse-Firmen kategorisieren Addresses nach Script-Typ, identifizieren Multisig-Wallets und leiten Verwahrungsstrukturen ab.

FROST beseitigt diesen Fingerabdruck. Ein 2-of-3-Custody-Wallet, ein 5-of-7-Unternehmens-Treasury und ein Single-Key-Privat-Wallet erzeugen alle identische On-Chain-Outputs. Jeder wird mit einem 32-Byte-Public-Key und einer 64-Byte-Signature innerhalb eines Taproot-Key-Path ausgegeben.

Diese Ununterscheidbarkeit kommt allen Taproot-Nutzern zugute. Je größer die Menge an Transactions, die gleich aussehen, desto schwerer ist es, eine einzelne Transaction herauszufiltern. Threshold-Signature-Nutzer verschwinden in der Masse der Single-Key-Sender, und Single-Key-Sender erhalten plausible Abstreitbarkeit, dass sie Threshold-Unterzeichner sein könnten.
