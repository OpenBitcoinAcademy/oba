## Ark: Off-Chain UTXOs

Ark ist ein Layer-2-Protokoll, das günstige Bitcoin-Transactions ohne Channels ermöglicht. Ein Ark Service Provider (ASP) koordiniert Runden, in denen Nutzer Virtual UTXOs (VTXOs) austauschen: gültige Bitcoin-Transactions, die off-chain gehalten werden.

Jeder VTXO ist ein Blatt in einem Baum von Transactions, der in einem einzelnen On-Chain Shared UTXO verwurzelt ist. Nutzer können ihren VTXO jederzeit senden, um Mittel on-chain zu beanspruchen (unilateraler Exit). Im Normalbetrieb bündelt der ASP Tausende Transfers in einer On-Chain-Abwicklung.

Ark erfordert kein Öffnen von Channels und kein Sperren von Liquidität im Voraus. Nutzer erhalten VTXOs, indem sie alte VTXOs in periodischen Runden aufgeben. Atomic Swaps über Connector Outputs stellen sicher, dass beim Austausch keine Mittel verloren gehen.
