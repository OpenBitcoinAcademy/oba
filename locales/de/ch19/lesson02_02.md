## Client-Side Validation

Bei RGB validieren nur die Vertragsparteien die Zustandsübergänge. Ein Token-Emittent und die Token-Inhaber verifizieren die Transferhistorie. Kein anderer Node im Netzwerk muss diese Daten verarbeiten oder speichern.

Vergleiche das mit Systemen, bei denen jeder Node jeden Vertrag validiert (Ethereums Modell). Client-Side Validation skaliert unbegrenzt: Mehr Verträge erhöhen nicht die Last auf Nodes, die nicht beteiligt sind. Privatsphäre ist inhärent: Vertragsdaten sind nur für Teilnehmer sichtbar.

Die Kosten: Teilnehmer müssen die vollständige Historie ihrer Assets speichern und verifizieren. Wenn die Historie verloren geht, kann das Asset nicht verifiziert werden. RGB nutzt Commitment-Ketten und Proofs, um diese Verifikation effizient zu gestalten.
