## Block-Validierungsregeln

Wenn eine Node einen neuen Block empfängt, prüft sie Dutzende Kriterien, bevor sie ihn akzeptiert. Der Hash des Block-Headers muss unter dem aktuellen Target liegen. Der Zeitstempel muss innerhalb akzeptabler Grenzen liegen. Die erste Transaction muss eine Coinbase sein, und der Coinbase-Output darf die erlaubte Subsidy plus Fees nicht übersteigen. Jede andere Transaction muss gemäß den Script-Regeln gültig sein.

Die Merkle Root des Blocks muss mit dem Hash-Baum aller enthaltenen Transactions übereinstimmen. Die Block-Größe darf das Konsenslimit nicht überschreiten. Der Block muss einen gültigen Parent-Block referenzieren, den die Node bereits hat.

Wenn eine Prüfung fehlschlägt, lehnt die Node den Block ab und trennt möglicherweise die Verbindung zum Peer, der ihn gesendet hat. Es gibt kein Einspruchsverfahren. Ein Block ist entweder gültig oder nicht. Diese strenge Validierung verhindert, dass Miner Bitcoin aus dem Nichts erzeugen oder Coins ausgeben, die ihnen nicht gehören.
