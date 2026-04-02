## Per-Input-Felder

Jede Input Map beschreibt einen Input der Transaction. Die kritischen Felder, die ein Signer braucht, sind die UTXO-Daten, der Ableitungspfad des Signatur-Keys und alle Scripts, die zur Konstruktion der Witness nötig sind.

WITNESS_UTXO (Key-Typ 0x01) speichert den auszugebenden Output: den Betrag in Satoshis und den scriptPubKey. Für SegWit-Inputs reicht das zum Signieren, weil der Sighash-Algorithmus den Betrag direkt committet. NON_WITNESS_UTXO (Key-Typ 0x00) speichert die gesamte vorherige Transaction, nötig für Legacy-Inputs, bei denen der Signer den Betrag durch Nachschlagen des Outputs in der vollständigen Transaction verifizieren muss.

BIP32_DERIVATION (Key-Typ 0x06) bildet einen Public Key auf seinen BIP 32 Ableitungspfad und Master-Key-Fingerprint ab. Der Signer gleicht den Fingerprint mit seinem eigenen Master Key ab, leitet den Private Key am angegebenen Pfad ab und signiert. PARTIAL_SIG (Key-Typ 0x02) speichert eine Signature, geschlüsselt nach dem Public Key, der sie erzeugt hat. SIGHASH_TYPE (Key-Typ 0x03) gibt an, welches Sighash-Flag der Signer verwenden soll.

Für P2SH-Inputs trägt REDEEM_SCRIPT (Key-Typ 0x04) das Redeem Script. Für P2WSH-Inputs trägt WITNESS_SCRIPT (Key-Typ 0x05) das Witness Script. Der Signer braucht diese, um den korrekten Sighash zu berechnen.
