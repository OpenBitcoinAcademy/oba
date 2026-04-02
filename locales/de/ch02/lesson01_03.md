## Transactions, Mining und die Blockchain

Eine Transaction ist eine Datenstruktur, die Bitcoin von einem Besitzer zum nächsten überträgt. Inputs referenzieren vorherige Outputs, die ausgegeben werden. Outputs erzeugen neue ausgebbare Beträge, die an die Keys der Empfänger gebunden sind.

Mining ist der Prozess, der Transactions zur Blockchain hinzufügt. Miner konkurrieren darum, ein rechnerisches Rätsel (Proof of Work) zu lösen, indem sie Block-Kandidaten hashen, bis das Ergebnis unter einen Zielwert fällt. Der erste Miner mit einer gültigen Lösung verbreitet den Block. Andere Nodes verifizieren ihn und akzeptieren ihn.

Die Blockchain ist die Kette aller gültigen Blocks, verknüpft durch Hashes. Jeder Block referenziert den vorherigen Block. Eine Änderung an einem vergangenen Block würde erfordern, den Proof of Work für diesen Block und jeden Block danach neu zu berechnen.
