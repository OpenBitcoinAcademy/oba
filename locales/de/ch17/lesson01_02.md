## Von n-of-n zu t-of-n

MuSig2 hat eine Einschränkung: Alle n Teilnehmer müssen signieren. Jeder Schlüsselinhaber muss zum Signierzeitpunkt online und kooperationsbereit sein. Wenn ein Teilnehmer seinen Key verliert oder offline geht, sind die Mittel blockiert. MuSig2 ist ein reines n-of-n-Schema.

FROST (Flexible Round-Optimized Schnorr Threshold signatures) löst dieses Problem. FROST ist ein Threshold-Signature-Schema: Beliebige t von n Teilnehmern können eine gültige Signature erzeugen. Ein 3-of-5-FROST-Setup bedeutet, dass beliebige drei der fünf Schlüsselinhaber signieren können. Die anderen zwei können offline, nicht erreichbar oder sogar dauerhaft verloren sein.

Die Signature, die FROST erzeugt, ist eine Standard-64-Byte-Schnorr-Signature. On-chain ist sie nicht von einem Single-Key-Spend zu unterscheiden. Die Threshold Policy, die Anzahl der Unterzeichner und die einzelnen Public Keys sind alle verborgen. Ein FROST-Spend sieht identisch aus wie ein Key-Path-Spend von einer einzelnen Taproot-Address.
