## Die drei Maps

Eine PSBT enthält drei Arten von Key-Value-Maps: eine globale Map, eine Map pro Input und eine Map pro Output. Jede Map speichert typisierte Key-Value-Paare, abgeschlossen durch ein 0x00-Byte. Der Key-Typ bestimmt die Bedeutung des Eintrags. Key-Payload und Value-Payload tragen die Daten.

Die globale Map enthält Daten, die für die gesamte Transaction gelten. In PSBTv0 ist der wichtigste globale Eintrag UNSIGNED_TX (Key-Typ 0x00), der die vollständige unsignierte Transaction enthält. In PSBTv2 enthält die globale Map stattdessen TX_VERSION, FALLBACK_LOCKTIME, INPUT_COUNT und OUTPUT_COUNT als separate Felder. Die unsignierte Transaction wird aus Per-Input- und Per-Output-Maps rekonstruiert statt als Ganzes gespeichert.

Globale Einträge enthalten auch XPUB (Key-Typ 0x01), der einen Extended Public Key auf seinen BIP 32 Ableitungspfad abbildet. Das erlaubt einem Signer zu verifizieren, dass Change Outputs von derselben Wallet-Root abgeleitet werden, ohne Zugriff auf den vollständigen Wallet Descriptor zu benötigen.
