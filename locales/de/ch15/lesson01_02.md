## Ressourcenlimits und versteckte Fehler

Bitcoin-Konsens setzt Ressourcenlimits für Scripts durch: eine Maximalgröße von 10.000 Bytes bei Legacy, ein Limit von 201 Non-Push-Opcodes und ein Sigops-Budget. Ein Script, das ein Limit überschreitet, ist ungültig. Die Transaction, die es enthält, wird nie bestätigt.

Bei handgeschriebenen Scripts ist der Ressourcenverbrauch schwer vorherzusagen. Jeder Zweig eines OP_IF-Baums trägt unterschiedlich zur Opcode-Zählung bei. Verschachtelte Bedingungen vervielfachen die Komplexität. Ein Script könnte in einem Zweig funktionieren und in einem anderen das Opcode-Limit überschreiten, abhängig davon, welchen Ausführungspfad ein Absender wählt.

Diese Fehler sind unsichtbar. Das Script sieht in einem Texteditor korrekt aus. Es besteht möglicherweise Unit-Tests für einen Ausgabepfad. Der fehlschlagende Pfad wird erst entdeckt, wenn jemand ihn im Netzwerk verwenden will und die Transaction abgelehnt wird.
