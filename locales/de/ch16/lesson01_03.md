## Warum Hardware Wallets das brauchten

Hardware Wallets haben eine spezifische Einschränkung: sie können die Blockchain nicht abfragen. Sie empfangen Daten, verifizieren was sie können, signieren und geben das Ergebnis zurück. Ohne Standardformat musste jeder Hardware-Wallet-Hersteller individuelle Integrationen für jede Software-Wallet schreiben. Ein neues Hardware Wallet hinzuzufügen bedeutete, jeden Coordinator zu patchen. Einen neuen Coordinator hinzuzufügen bedeutete, jedes Hardware Wallet zu patchen. Die kombinatorische Explosion war nicht tragbar.

PSBT löste das durch die Definition eines Formats, das alle Informationen trägt, die ein Signer braucht. Den auszugebenden UTXO, den Ableitungspfad für den Key, den zu verwendenden Sighash-Typ, das Redeem Script für P2SH, das Witness Script für P2WSH. Ein Hardware Wallet empfängt eine PSBT, liest die benötigten Felder, signiert, schreibt seine partielle Signature zurück in die PSBT und gibt sie zurück. Kein proprietäres Protokoll. Keine Formatverhandlung.

Das Ökosystem konvergierte schnell. Coldcard, Trezor, Ledger, BitBox und Jade sprechen alle PSBT. Software-Coordinatoren wie Sparrow, Specter und Bitcoin Core erzeugen und konsumieren alle PSBTs. Ein Multisig-Quorum kann Hardware verschiedener Hersteller verwenden, ohne Kompatibilitätssorgen.
