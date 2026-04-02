## Bitcoin-Systeme sicher bauen

Entwickler, die auf Bitcoin aufbauen, stehen vor einer besonderen Herausforderung: Bugs können Geld kosten. Ein Fehler bei der Key-Erzeugung, Transaction-Erstellung oder Signature-Validierung kann dazu führen, dass Mittel dauerhaft gestohlen werden oder verloren gehen.

Das dezentrale Konsensmodell kennt keine Instanz, die eine fehlerhafte Transaction rückgängig machen kann. Code, der mit Private Keys umgeht, muss diese als die sensibelsten Daten im System behandeln. Keys sollten aus hochwertigen Entropiequellen erzeugt, verschlüsselt gespeichert und nach Gebrauch aus dem Speicher gelöscht werden.

Jede Komponente, die Keys berührt oder Transactions konstruiert, sollte auditiert, gegen bekannte Testvektoren geprüft und einer adversariellen Überprüfung unterzogen werden.
