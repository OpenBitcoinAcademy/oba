## Weight Units

Segwit führte Weight Units ein, die das bytebasierte Größenmaß ersetzen. Jedes Byte an Nicht-Witness-Daten kostet 4 Weight Units. Jedes Byte an Witness-Daten (Signatures) kostet 1 Weight Unit.

Dieser Rabatt schafft einen Anreiz für die Nutzung von Segwit-Transactions, die Signatures in der Witness-Struktur speichern. Eine Segwit-Transaction verbraucht weniger Gewicht als eine Legacy-Transaction mit derselben Anzahl an Inputs und Outputs.

Virtual Bytes (vbytes) rechnen Gewicht in ein Byte-Äquivalent um: vbytes = weight / 4. Eine Transaction mit 600 Weight Units hat 150 vbytes. Fee-Raten in sat/vB berücksichtigen den Segwit-Rabatt automatisch.
