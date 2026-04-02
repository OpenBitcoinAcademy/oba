Was passiert, wenn ein FROST-Unterzeichner dieselbe Nonce in zwei verschiedenen Signiersitzungen wiederverwendet?

- Die finale Signature wird ungültig, aber keine geheimen Daten werden preisgegeben
- Ein Angreifer kann den Secret Key Share dieses Unterzeichners aus den zwei Signature Shares extrahieren
- Der Coordinator erkennt die Wiederverwendung und bricht beide Sitzungen ab
