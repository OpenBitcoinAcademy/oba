## Bitcoin Script von Hand schreiben

Bitcoin Script ist eine Stack-basierte Sprache. Jede Ausgabebedingung ist eine Folge von Opcodes: einen Key auf den Stack legen, eine Signature prüfen, einen Hash verifizieren, einen Timelock durchsetzen. Für eine Einzel-Key-Ausgabe ist das Script kurz und gut verstanden.

Kombiniere Bedingungen, und die Schwierigkeit steigt schnell. Ein 2-of-3 Multisig mit Timelock-Rückfall erfordert sorgfältige Anordnung von OP_IF-Verzweigungen, OP_CHECKMULTISIG und OP_CHECKSEQUENCEVERIFY. Ein falsch platzierter Opcode ändert die Ausgabebedingungen vollständig. Ein doppelter Key fällt möglicherweise erst auf, wenn Mittel gesperrt sind.

Kein Werkzeug kann ein beliebiges Script analysieren und beantworten: "Wer kann das ausgeben, unter welchen Bedingungen?" Der Sprache fehlt Struktur. Verschiedene Opcode-Sequenzen können dieselbe Policy kodieren, und es gibt keine allgemeine Methode, sie zu vergleichen.
