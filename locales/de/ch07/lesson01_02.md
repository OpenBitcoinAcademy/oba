## Unlocking Scripts

Um einen gesperrten Output auszugeben, liefert der Absender ein Unlocking Script, genannt scriptSig. Für P2PKH enthält das Unlocking Script eine digitale Signature und den Public Key des Absenders.

Bitcoin führt diese Scripts auf einer Stack-Maschine aus. Zuerst läuft das Unlocking Script und legt Daten auf den Stack. Dann wird der Stack kopiert, und das Locking Script läuft gegen diesen kopierten Stack. Die beiden Scripts werden nie zu einem kombiniert. Diese Trennung wurde 2010 eingeführt, um eine Sicherheitslücke zu schließen.

Wenn das Locking Script mit einem wahren Wert oben auf dem Stack endet, ist die Ausgabe gültig. Wenn der Stack leer ist oder der oberste Wert null ist, schlägt die Ausgabe fehl.
