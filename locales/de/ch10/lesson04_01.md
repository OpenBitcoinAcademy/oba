## SPV-Clients

Nicht jedes Gerät kann die vollständige Blockchain speichern und validieren. Mobiltelefone, eingebettete Geräte und Hardware Wallets haben begrenzten Speicher und begrenzte Rechenleistung. Diese Geräte nutzen **Simplified Payment Verification** (SPV) Clients.

Ein SPV-Client lädt nur Block-Header herunter, keine vollständigen Blocks. Die Header-Chain ist klein: etwa 60 MB für die gesamte Geschichte von Bitcoin. Der Client kann verifizieren, dass ein Block-Header das Proof-of-Work-Target erfüllt, und damit bestätigen, dass ein Miner echte Energie aufgewendet hat.

Um zu prüfen, ob eine Transaction bestätigt ist, fordert der SPV-Client einen **Merkle Proof** von einer Full Node an. Der Proof zeigt, dass der Hash der Transaction in der Merkle Root eines Blocks enthalten ist. Der Client vertraut darauf, dass Full Nodes die Transactions des Blocks validiert haben, weil das Fälschen eines gültigen Proof-of-Work-Headers unerschwinglich teuer ist.

SPV gibt Lightweight-Clients ein vernünftiges Maß an Sicherheit, ohne die vollständige Blockchain herunterzuladen. Der Kompromiss: Ein SPV-Client kann nicht erkennen, ob ein Block eine ungültige Transaction enthält. Er vertraut auf das ökonomische Gewicht des Proof of Work, statt jede Regel selbst zu prüfen.
