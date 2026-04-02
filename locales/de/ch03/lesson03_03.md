## Das Datenverzeichnis

Bitcoin Core speichert seine Daten in einem plattformspezifischen Verzeichnis: `~/.bitcoin` auf Linux, `~/Library/Application Support/Bitcoin` auf macOS, `%APPDATA%\Bitcoin` auf Windows.

Das Unterverzeichnis `blocks/` enthält die rohen Blockdaten-Dateien. Das Verzeichnis `chainstate/` enthält die UTXO-Datenbank, einen LevelDB-Speicher aller unverbrauchten Outputs. Die Datei `mempool.dat` speichert den Mempool über Neustarts hinweg.

Die Konfigurationsdatei `bitcoin.conf` steuert Netzwerkeinstellungen, Ressourcenlimits, RPC-Authentifizierung und Feature-Flags. Einstellungen wie `prune=550` aktivieren den Pruned-Modus und `txindex=1` erstellt einen Index aller Transactions für schnellere Abfragen.
