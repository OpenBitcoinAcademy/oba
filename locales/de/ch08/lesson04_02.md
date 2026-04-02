## Die Zufallsanforderung

Sowohl ECDSA als auch Schnorr erfordern eine Zufallszahl (Nonce) bei der Signature-Erstellung. Die Sicherheit des gesamten Systems hängt von dieser Zufälligkeit ab.

Wenn dieselbe Nonce verwendet wird, um zwei verschiedene Nachrichten mit demselben Private Key zu signieren, kann der Private Key aus den beiden Signatures berechnet werden. Das ist kein theoretisches Risiko. 2013 verursachte ein Fehler im Android-Zufallszahlengenerator, dass mehrere Bitcoin-Wallets Nonces wiederverwendeten, und Guthaben wurden gestohlen.

Moderne Implementierungen verwenden deterministische Nonces (RFC 6979 für ECDSA, BIP340 für Schnorr), die die Nonce aus dem Private Key und dem Nachrichten-Hash ableiten. Das beseitigt die Abhängigkeit von einem Zufallszahlengenerator und verhindert Nonce-Wiederverwendung.
