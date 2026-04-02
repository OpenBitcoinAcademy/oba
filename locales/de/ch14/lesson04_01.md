## Eine neue Script-Sprache

Tapscript (BIP 342) definiert die Validierungsregeln für Scripts, die in Taproot Script Path Spends ausgeführt werden. Es ähnelt dem bisherigen Bitcoin Script, mit gezielten Verbesserungen.

Alle Signature-prüfenden Opcodes (OP_CHECKSIG, OP_CHECKSIGVERIFY) verwenden Schnorr Signatures statt ECDSA. Ein neuer Opcode, OP_CHECKSIGADD, ersetzt das bisherige OP_CHECKMULTISIG. Statt mehrere Signatures gegen mehrere Keys in einer Operation zu prüfen, prüft OP_CHECKSIGADD jeweils eine Signature und akkumuliert einen Zähler. Das ermöglicht Batch-Verifikation: der Verifizierer sammelt alle Signatures und verifiziert sie in einer einzigen Operation, schneller als jede einzeln zu prüfen.
