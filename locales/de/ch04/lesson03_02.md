## Die Ableitungsschritte

Das Erstellen einer Bitcoin Address aus einem Public Key erfordert mehrere Hashing-Schritte.

Zuerst wird der Public Key mit SHA-256 gehasht. Dann wird dieses Ergebnis mit RIPEMD-160 gehasht. Das erzeugt einen 20-Byte-Hash, den Public Key Hash oder HASH160.

Als Nächstes wird ein Versionsbyte vorangestellt. Dieses Byte identifiziert das Netzwerk (Mainnet vs. Testnet) und den Address-Typ.

Zum Schluss wird das Ergebnis kodiert. Legacy-Addresses verwenden Base58Check-Kodierung, die eine 4-Byte-Prüfsumme anfügt und in einem lesbaren Alphabet kodiert, das verwechselbare Zeichen wie 0/O und l/1 vermeidet. Neuere SegWit-Addresses verwenden Bech32-Kodierung: nur Kleinbuchstaben und stärkere Fehlererkennung.
