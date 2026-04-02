## Der Handshake

Zwei Nodes stellen eine Verbindung mit einem Version/Verack-Handshake über TCP-Port 8333 (den Standard-Bitcoin-Netzwerkport) her.

Die verbindende Node sendet eine `version`-Nachricht mit ihrer Protokollversion, Block-Höhe, aktueller Zeit und den Diensten, die sie unterstützt. Die empfangende Node antwortet mit ihrer eigenen `version`-Nachricht. Beide Nodes senden dann ein `verack` (Version Acknowledgement) zur Bestätigung.

Nach Abschluss des Handshakes können die Nodes Daten austauschen. Wenn die Protokollversionen inkompatibel sind, wird die Verbindung getrennt.

Nodes teilen bekannte Peer-Adressen über `addr`- und `getaddr`-Nachrichten. Wenn eine Node neue Adressen erfährt, speichert sie diese und verbindet sich möglicherweise später damit. Dieses Gossip-Protokoll lässt das Netzwerk wachsen und sich selbst reparieren, ohne ein zentrales Verzeichnis.
