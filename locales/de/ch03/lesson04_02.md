## Die Implementierungslandschaft

Mehrere Teams pflegen unabhängige Bitcoin-Implementierungen in verschiedenen Programmiersprachen.

**btcd** (Go) ist eine von Grund auf neu geschriebene Full-Node-Implementierung. Sie betreibt die Infrastruktur mehrerer Bitcoin-Unternehmen und bildet die Basis von lnd, dem Lightning Network.

**bcoin** (JavaScript) ist eine modulare Full-Node-Implementierung für das Node.js-Ökosystem, mit eingebautem Wallet und HTTP-API.

**Bitcoin Knots** ist ein Fork von Bitcoin Core, gepflegt von Luke Dashjr. Er enthält zusätzliche Konfigurationsoptionen und strengere Policy-Standardwerte.

**rust-bitcoin** ist eine Bibliothek für die Arbeit mit Bitcoin-Datenstrukturen in Rust. Sie bietet Serialisierung, Parsing und Scripting-Werkzeuge, ohne eine Full Node zu betreiben.

**libbitcoin** (C++) ist ein unabhängiges Toolkit zum Erstellen von Bitcoin-Anwendungen, einschließlich einer Full-Node-Implementierung namens libbitcoin-node.

Weitere Implementierungen existieren in Python, Java, Scala und C#. Jede bringt eine andere Entwicklergemeinschaft ins Bitcoin-Ökosystem.
