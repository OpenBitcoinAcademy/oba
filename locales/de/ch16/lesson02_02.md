## Updater und Signer

Der Updater fügt die Metadaten hinzu, die Signer brauchen. Für jeden Input hängt der Updater den auszugebenden UTXO an (entweder die vollständige vorherige Transaction oder den spezifischen Witness UTXO), den BIP 32 Ableitungspfad für den Signatur-Key, das Redeem Script bei P2SH-Inputs und das Witness Script bei P2WSH-Inputs. Für jeden Output kann der Updater BIP 32 Ableitungspfade anhängen, damit der Signer verifizieren kann, dass Change Outputs zum selben Wallet gehören.

Der Signer liest die PSBT, identifiziert die Inputs, die er signieren kann, und erzeugt partielle Signatures. Für jeden Input, den er signiert, schreibt er einen PARTIAL_SIG-Eintrag, geschlüsselt nach dem Public Key. Der Signer ändert die Transaction selbst nicht. Er finalisiert keine Scripts. Er fügt seine Signature hinzu und reicht die PSBT an den nächsten Beteiligten weiter.

Ein Signer muss die UTXO-Daten vor dem Signieren verifizieren. Behauptet die PSBT, ein Input gibt 0,5 BTC aus, der tatsächliche UTXO enthält aber 1,0 BTC, würde der Signer unwissentlich 0,5 BTC an Gebühren spenden. Hardware Wallets vergleichen UTXO-Beträge mit Output-Beträgen und warnen den Nutzer, wenn die Gebühr ungewöhnlich hoch erscheint.
