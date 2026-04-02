## Konsensrisiken

Anwendungen, die Transactions validieren, müssen die Konsensregeln exakt umsetzen. Ein subtiler Unterschied in der Behandlung eines Grenzfalls durch zwei Implementierungen kann dazu führen, dass sie sich über die Gültigkeit eines Blocks uneinig sind und aus ihrer Perspektive das Netzwerk spalten.

Der sicherste Ansatz für Anwendungen: delegiere die Konsensvalidierung an einen Full Node (Bitcoin Core oder gleichwertig) und verwende dessen API für Blockchain-Abfragen. Implementiere keine Konsensregeln in Anwendungscode neu, es sei denn, du bist bereit, jeden Grenzfall der Referenzimplementierung abzudecken.

Für Wallet-Entwickler: verwende bewährte Bibliotheken für Key-Ableitung, Address-Erzeugung und Transaction-Signierung. Bevorzuge kampferprobte Open-Source-Implementierungen gegenüber eigenem Code.
