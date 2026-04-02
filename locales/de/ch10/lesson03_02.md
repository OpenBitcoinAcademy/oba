## Compact Block Relay

Definiert in **BIP 152**, reduziert Compact Block Relay die Bandbreite, die nötig ist, um einen neuen Block zu verbreiten. Die zentrale Erkenntnis: Die meisten Transactions in einem neuen Block befinden sich bereits im Mempool der empfangenden Node.

Statt den vollständigen Block zu senden, schickt die ankündigende Node eine `cmpctblock`-Nachricht. Diese enthält den Block-Header, eine Liste kurzer Transaction-IDs und die Coinbase-Transaction. Kurze Transaction-IDs sind 6-Byte-gekürzte Hashes, mit denen der Empfänger Transactions zuordnen kann, die er bereits hat.

Die empfangende Node rekonstruiert den Block aus ihrem eigenen Mempool mithilfe der kurzen IDs. Wenn Transactions fehlen, fordert sie diese mit einer `getblocktxn`-Nachricht an und erhält sie in einer `blocktxn`-Antwort.

BIP 152 definiert zwei Modi. Im **Low-Bandwidth-Modus** kündigt eine Node den Block zuerst mit einer `inv`-Nachricht an. Der Peer fordert den Compact Block nur bei Interesse an. Im **High-Bandwidth-Modus** sendet die Node den Compact Block sofort, ohne auf eine Anfrage zu warten. Miner und gut vernetzte Nodes nutzen typischerweise den High-Bandwidth-Modus, um die Latenz zu minimieren.

Compact Block Relay reduziert die Verbreitungsbandbreite um 90 % oder mehr bei einem typischen Block. Schnellere Verbreitung gibt kleineren Minern fairere Chancen und stärkt die Dezentralisierung.
