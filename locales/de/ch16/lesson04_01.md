## Der Watch-Only-Workflow

Ein Hardware-Wallet-Signierworkflow beginnt mit einem Watch-Only-Wallet auf einem vernetzten Computer. Das Watch-Only-Wallet hält die Extended Public Keys (xpubs) aller Signer, aber keine Private Keys. Es kann Addresses erzeugen, Kontostände verfolgen und Transactions erstellen, aber nicht signieren.

Wenn der Nutzer ausgeben will, erstellt das Watch-Only-Wallet eine PSBT. Es füllt die unsignierte Transaction aus, hängt WITNESS_UTXO-Daten für jeden Input an und schreibt BIP32_DERIVATION-Einträge, damit das Hardware Wallet weiß, welche Keys zu verwenden sind. Die PSBT wird als Datei auf eine SD-Karte exportiert oder als animierter QR-Code kodiert.

Das Hardware Wallet empfängt die PSBT, parst Inputs und Outputs und zeigt dem Nutzer eine Zusammenfassung: welche Addresses Mittel empfangen, wie viel an jede geht und wie viel an Gebühren gezahlt wird. Der Nutzer bestätigt auf dem Gerät. Das Hardware Wallet leitet die Private Keys aus seinem Seed über die BIP 32 Pfade in der PSBT ab, signiert jeden möglichen Input, schreibt PARTIAL_SIG-Einträge und exportiert die aktualisierte PSBT zurück per SD-Karte oder QR-Code.

Das Watch-Only-Wallet importiert die signierte PSBT. Sind alle nötigen Signatures vorhanden, finalisiert und extrahiert es die rohe Transaction und sendet sie. Werden weitere Signatures benötigt (wie bei einem Multisig-Setup), reicht es die PSBT an den nächsten Signer weiter.
