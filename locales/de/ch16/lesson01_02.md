## Was eine PSBT ist

Eine Partially Signed Bitcoin Transaction (PSBT) ist ein Binärformat, definiert durch BIP 174. Es verpackt eine unsignierte Transaction zusammen mit den Metadaten, die jeder Beteiligte für seine Aufgabe braucht. Creators hängen UTXO-Daten an. Updater fügen Ableitungspfade hinzu. Signer fügen Signatures hinzu. Das Format trägt alles Nötige, damit jede Rolle unabhängig handeln kann.

Das Binärformat beginnt mit einem Fünf-Byte-Magic: `0x70736274FF`. Die ersten vier Bytes buchstabieren "psbt" in ASCII. Das fünfte Byte, `0xFF`, markiert den Separator. Jedes Werkzeug, das einen Blob empfängt, kann diese fünf Bytes prüfen, um zu bestätigen, dass es sich um eine PSBT handelt, bevor es weiter parst.

Nach dem Magic folgt eine Sequenz von Key-Value-Maps. Jeder Map-Eintrag hat einen Key-Typ, einen Key-Payload und einen Value-Payload. Ein Null-Byte (0x00) beendet jede Map. Die erste Map ist die globale Map. Nachfolgende Maps wechseln zwischen Per-Input- und Per-Output-Maps, je eine für jeden Input und Output der Transaction.
