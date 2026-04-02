## Archival und Pruned Nodes

Eine Full Node, die jeden Block seit dem Genesis Block aufbewahrt, wird **Archival Node** genannt. Sie kann historische Block-Daten an jeden Peer liefern, der sie anfragt. Neue Nodes, die dem Netzwerk beitreten, sind auf Archival Nodes angewiesen, um die gesamte Blockchain herunterzuladen.

Eine Archival Node zu betreiben erfordert Hunderte Gigabyte Speicherplatz. Stand 2024 übersteigt die Blockchain 600 GB. Nicht jeder Full-Node-Betreiber kann sich diesen Speicher leisten.

Eine **Pruned Node** validiert jeden Block, verwirft aber alte Block-Daten nach der Verarbeitung. Sie behält nur das UTXO-Set und die neuesten Blocks. Eine Pruned Node setzt jede Konsensregel durch. Sie basiert nicht auf Vertrauen. Der Kompromiss: Sie kann keine historischen Blocks an andere Peers liefern.

Sowohl Archival als auch Pruned Nodes sind Full Nodes. Beide validieren alles. Der Unterschied liegt darin, ob sie alte Blocks für den Download durch andere speichern.
