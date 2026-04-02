## Die JSON-RPC-Schnittstelle

Bitcoin Core stellt eine JSON-RPC-Schnittstelle bereit, über die Programme Blockchain-Daten abfragen und Transactions einreichen können. Das Kommandozeilenwerkzeug `bitcoin-cli` sendet Anfragen an diese Schnittstelle.

Jedes Datenelement in der Blockchain ist abfragbar: Block-Header, vollständige Blocks, einzelne Transactions, Kontostände von Addresses, Mempool-Inhalte, Peer-Verbindungen und Netzwerkstatistiken.

Wallets, Block-Explorer, Zahlungsdienstleister und Forschungswerkzeuge kommunizieren alle über diese API mit Bitcoin Core. Die Schnittstelle ist die Brücke zwischen den rohen Blockchain-Daten und den Anwendungen, die sie den Nutzern präsentieren.
