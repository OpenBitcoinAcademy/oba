## Compact Block Filters

**BIP 157** und **BIP 158** definieren einen besseren Ansatz namens Compact Block Filters. Statt dass der Client seinen Filter an den Server sendet, erstellt der Server einen Filter für jeden Block und sendet ihn an den Client.

Jeder Filter ist eine kompakte Darstellung aller Addresses und Scripts in einem Block. Der Client lädt den Filter herunter und prüft lokal, ob eine seiner Addresses vorkommt. Bei einem Treffer lädt der Client den vollständigen Block herunter, um die relevanten Transactions zu extrahieren.

Der Datenschutzvorteil: Der Server erfährt nie, welche Addresses den Client interessieren. Der Server sendet denselben Filter an jeden Client. Der Client erledigt den gesamten Abgleich lokal.

Die Bandbreitenkosten sind höher als bei Bloom Filtern, weil der Client einen Filter für jeden Block herunterlädt. Die Filter sind aber klein (durchschnittlich etwa 20 KB pro Block) und können gegen ein Hash-Commitment in der Block-Header-Chain verifiziert werden. Der Client muss dem Server nicht vertrauen, dass er korrekte Filter liefert.

Bitcoin über **Tor** zu betreiben fügt eine weitere Datenschutzschicht hinzu. Tor verbirgt die IP-Adresse des Clients vor den Nodes, mit denen er sich verbindet. Kombiniert mit Compact Block Filters kann ein Lightweight-Client seinen Kontostand prüfen, ohne seine Identität oder seine Addresses einem Peer zu offenbaren.
