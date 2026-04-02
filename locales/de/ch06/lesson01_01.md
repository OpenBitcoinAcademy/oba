## Wie eine Transaction aussieht

Eine Bitcoin-Transaction ist eine Datenstruktur, die Bitcoin von einem Besitzer auf einen anderen überträgt. Sie enthält Inputs (die auf vorherige Outputs verweisen, um sie auszugeben), Outputs (die neue ausgebbare Beträge erzeugen) und Autorisierungsdaten (Signatures, die beweisen, dass der Absender die Keys kontrolliert).

Inputs geben an, woher die Bitcoin kommen. Jeder Input verweist auf einen Output einer früheren Transaction, der noch nicht ausgegeben wurde. Dieser unausgegebene Output heißt UTXO (Unspent Transaction Output).

Outputs geben an, wohin die Bitcoin gehen. Jeder Output legt einen Betrag in Satoshis und eine Sperrbedingung (ein Script) fest, die bestimmt, wer ihn ausgeben kann.

Eine Transaction verbraucht alte Outputs und erzeugt neue. Nichts wird "auf einem Konto gespeichert". Bitcoin verfolgt Eigentum über eine Kette von Outputs, jeder an einen bestimmten Key gebunden.
