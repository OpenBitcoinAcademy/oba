## Verbreitung im Netzwerk

Alice' Wallet signiert die Transaction mit ihrem Private Key und sendet sie ins Bitcoin-Netzwerk. Ihr Wallet verbindet sich mit mehreren Nodes, die die Transaction unabhängig verifizieren: korrektes Format, gültige Signature, unverbrauchte Inputs, Output-Werte übersteigen nicht die Input-Werte.

Jede Node, die die Transaction akzeptiert, fügt sie ihrem Mempool hinzu (einer lokalen Liste unbestätigter Transactions) und leitet sie an ihre Peers weiter. Innerhalb von Sekunden erreicht die Transaction Tausende Nodes im gesamten Netzwerk. Kein zentraler Server koordiniert das. Jede Node befolgt die gleichen Regeln unabhängig.

Bobs Wallet überwacht das Netzwerk auf Transactions an seine Address. Wenn es Alice' Transaction erkennt, wird sie als "unbestätigt" in Bobs Wallet angezeigt. Die Zahlung wurde gesendet, aber noch nicht in einem Block aufgezeichnet.
