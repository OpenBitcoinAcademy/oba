## Die Blockchain erkunden

Mit `bitcoin-cli` kannst du jeden Block über seine Höhe oder seinen Hash abrufen. Der Befehl `getblock` gibt die Block-Header-Felder, die Liste der Transaction-IDs und Metadaten wie die Anzahl der Bestätigungen zurück.

Der Befehl `getrawtransaction` gibt die rohen Hex-Daten einer Transaction oder ihre dekodierte JSON-Darstellung zurück und zeigt Inputs, Outputs, Scripts und Witness-Daten.

Der Befehl `getblockchaininfo` meldet die aktuelle Chain-Höhe, Schwierigkeit, Verifizierungsfortschritt und ob die Node noch synchronisiert. Diese Befehle verwandeln abstrakte Blockchain-Konzepte in konkrete, überprüfbare Daten.
