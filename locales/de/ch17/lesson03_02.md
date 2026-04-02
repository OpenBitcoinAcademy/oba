## Drei Schichten: SimplPedPop, EncPedPop, ChillDKG

Die Grundlage ist SimplPedPop. Jeder Teilnehmer erzeugt sein eigenes zufälliges Polynom vom Grad t-1 und sendet eine geheime Auswertung an jeden anderen Teilnehmer, zusammen mit einem Commitment auf die Koeffizienten seines Polynoms. Jeder Teilnehmer summiert die empfangenen Auswertungen, um seinen endgültigen Secret Share zu berechnen. Der Group Public Key wird aus der Summe aller Teilnehmer-Commitments auf ihre konstanten Terme abgeleitet. Keine Partei hält jemals das vollständige Secret.

SimplPedPop setzt einen sicheren Kanal zwischen jedem Teilnehmerpaar voraus. EncPedPop ergänzt dies, indem jeder Teilnehmer ein ephemeres Verschlüsselungs-Schlüsselpaar erzeugt und seine geheimen Auswertungen mit dem Public Key des Empfängers verschlüsselt. Nun funktioniert das Protokoll über einen öffentlichen Broadcast-Kanal, weil Lauscher die Secret Shares nicht entschlüsseln können.

ChillDKG umhüllt EncPedPop mit einer Session-Management-Schicht. Es führt einen Host Secret Key ein, den jeder Teilnehmer dauerhaft hält, ein gemeinsames Recovery-Dataset, das für alle Teilnehmer identisch ist, und ein Protokoll zur Erkennung und Behandlung von Fehlverhalten. Der Host Secret Key, kombiniert mit den gemeinsamen Recovery-Daten, ermöglicht einem Teilnehmer die Rekonstruktion seines Share, falls er sein Signiergerät verliert.
