## Target, Difficulty und die Nonce

Das Target ist eine 256-Bit-Zahl. Ein gültiger Block-Hash muss numerisch kleiner als das Target sein. Ein niedrigeres Target bedeutet weniger gültige Hashes, was das Finden eines gültigen Hashes schwieriger macht. Die Difficulty ist das Inverse des Targets: Wenn das Target sinkt, steigt die Difficulty.

Der Block-Header enthält ein 32-Bit-Nonce-Feld. Miner erhöhen diesen Wert mit jedem Hash-Versuch und durchlaufen alle $2^{32}$ (etwa 4,3 Milliarden) Möglichkeiten. Moderne Miner erschöpfen den Nonce-Raum in unter einer Sekunde.

Wenn der Nonce-Raum erschöpft ist, ändern Miner andere Felder, um neue Hash-Inputs zu erzeugen. Die gängigste Technik ändert das Extra-Nonce-Feld der Coinbase-Transaction. Das ändert die Merkle Root im Block-Header und gibt dem Miner einen frischen Satz von $2^{32}$ Nonces zum Ausprobieren.
