## P2PKH Schritt für Schritt

So läuft ein P2PKH-Script ab:

1. Die Signature und der Public Key werden auf den Stack gelegt (aus dem Unlocking Script).
2. OP_DUP dupliziert den Public Key auf dem Stack.
3. OP_HASH160 hasht das oberste Element (den duplizierten Public Key) mit SHA-256 und dann RIPEMD-160.
4. Der erwartete Public Key Hash wird auf den Stack gelegt (aus dem Locking Script).
5. OP_EQUALVERIFY nimmt die obersten zwei Elemente und prüft, ob sie gleich sind. Wenn sie sich unterscheiden, schlägt das Script fehl.
6. OP_CHECKSIG nimmt die Signature und den Public Key, verifiziert die Signature gegen die Transaction-Daten. Bei Gültigkeit legt es true auf den Stack.

Das Ergebnis: true auf dem Stack. Die Ausgabe ist autorisiert. Die Signature beweist, dass der Absender den Private Key kontrolliert, ohne ihn preiszugeben.
