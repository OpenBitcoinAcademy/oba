## Die Blockchain synchronisieren

Eine neue Node muss die gesamte Blockchain herunterladen und validieren, bevor sie aktuelle Transactions verifizieren kann. Dieser Vorgang heißt **Initial Block Download** (IBD).

Die Node sendet `getheaders`-Nachrichten an ihre Peers und fordert Block-Header in Stapeln an. Header sind klein (jeweils 80 Bytes) und kommen schnell an. Die Node nutzt sie, um die Header-Chain mit dem meisten kumulativen Proof of Work aufzubauen.

Sobald die Node die beste Header-Chain identifiziert hat, fordert sie vollständige Blocks über `getdata`-Nachrichten an. Sie lädt Blocks von mehreren Peers parallel herunter, um den Vorgang zu beschleunigen. Jeder Block wird bei Ankunft gegen die Konsensregeln validiert: Transaction-Signatures, Script-Ausführung, Gewichtslimits und das Proof-of-Work-Target.

IBD kann auf langsamer Hardware Stunden oder Tage dauern. Nach Abschluss wechselt die Node in den Normalbetrieb. Sie empfängt neue Blocks und Transactions, sobald diese gesendet werden, und validiert sie in Echtzeit.
