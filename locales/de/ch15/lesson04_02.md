## Multisig mit Zeit erweitern

Zeitgesperrte Policies können ihren Signer-Kreis im Lauf der Zeit erweitern. Ein Unternehmen könnte mit einer 2-of-2-Policy zwischen zwei Mitgründern beginnen. Nach einem Jahr wird ein dritter Key (gehalten vom Rechtsberater) aktiviert, und die Policy wird zu 2-of-3.

Der Miniscript-Ausdruck lautet: `or(multi(2,founder_a,founder_b),and(multi(2,founder_a,founder_b,counsel),older(52560)))`. Während des ersten Jahres können nur die zwei Gründer signieren. Nach 52.560 Blocks (ungefähr ein Jahr) können beliebige zwei der drei Keys signieren.

Die gesamte Policy lebt in einem einzigen UTXO. Keine On-chain-Transaction ist nötig, wenn der Timelock abläuft. Der zusätzliche Ausgabepfad wurde bei der Output-Erstellung committet. Der Key des Beraters erhält automatisch Ausgabeberechtigung, wenn die Blockhöhe den Schwellenwert überschreitet. Das entfernt die Notwendigkeit einer aktiven Key-Zeremonie am Übergangspunkt.

Miniscript macht diese Kompositionen auditierbar. Der Compiler kann jeden Ausgabepfad und seine Bedingungen auflisten. Ein Prüfer kann verifizieren, dass der Key des Beraters vor dem Timelock keine Ausgabeberechtigung hat, dass die Gründer jederzeit volle Kontrolle behalten, und dass die Witness-Größen innerhalb der Konsens-Limits bleiben.
