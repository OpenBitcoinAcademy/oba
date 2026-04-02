## Peers finden

Wenn eine Bitcoin-Node zum ersten Mal startet, kennt sie keine anderen Nodes. Sie muss mindestens einen Peer finden, um am Netzwerk teilzunehmen.

Bitcoin verwendet **DNS Seeds** zur anfänglichen Erkennung. DNS Seeds sind DNS-Server, die von Mitgliedern der Bitcoin-Community betrieben werden. Sie liefern die IP-Adressen bekannter, stabiler Full Nodes zurück. Bitcoin Core hat mehrere DNS Seeds fest in seinem Quellcode eingebaut.

Wenn DNS nicht verfügbar ist (blockiert oder zensiert), enthält Bitcoin Core auch eine Liste fest eingebauter IP-Adressen als letzten Ausweg. Diese Adressen werden mit jeder Software-Version aktualisiert.

Sobald eine Node sich mit ihrem ersten Peer verbindet, kann sie diesen Peer nach weiteren Adressen fragen. Die Node baut eine vielfältige Menge an Verbindungen auf, typischerweise 8 ausgehende und bis zu 125 eingehende.
