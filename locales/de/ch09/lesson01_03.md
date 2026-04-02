## Widersprüchliche Transactions

Alice kann zwei Transactions erstellen, die denselben Output ausgeben: eine an Bob und eine an Carol. Beide sind gültig signiert, aber nur eine kann in die Blockchain aufgenommen werden (der Output kann nur einmal ausgegeben werden).

Miner entscheiden, welche der widersprüchlichen Transactions sie aufnehmen. Standardmäßig folgen die meisten Miner einer "First Seen"-Regel und nehmen die Transaction auf, die sie zuerst erhalten haben. Das ist aber eine Richtlinie, keine Konsensregel. Ein Miner kann jede gültige Transaction aufnehmen.

Deshalb sollte Bob auf Bestätigungen warten, bevor er eine Zahlung als endgültig betrachtet. Die Fee, die Alice zahlt, gibt Minern einen Anreiz, ihre Transaction schnell aufzunehmen. Das verkleinert das Zeitfenster für widersprüchliche Transactions.
