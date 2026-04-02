## Multisig-Weiterreichung und Transport

In einer Multisig-Konfiguration wandert die PSBT der Reihe nach durch jeden Signer oder wird parallel an alle Signer verteilt. Betrachte ein 2-of-3 Multisig: der Coordinator erstellt die PSBT und sendet sie gleichzeitig an Signer A und Signer B. Jeder Signer fügt seine partielle Signature hinzu und gibt die PSBT zurück. Der Coordinator kombiniert beide PSBTs, finalisiert und sendet.

Der Transportmechanismus ist durch BIP 174 nicht festgelegt. PSBTs können auf SD-Karten, als QR-Codes (im UR-Encoding nach BCR-2020-005), über NFC, per E-Mail, über einen Filesharing-Dienst oder über jeden anderen Kanal übertragen werden. Das Format ist in sich geschlossen. Jeder Signer bekommt alles Nötige aus der PSBT selbst, ohne Seitenkanal.

PSBT ist Klartext, nicht verschlüsselt. Wer eine PSBT abfängt, kann die Transaction-Beträge, die beteiligten Addresses und die bisher gesammelten partiellen Signatures sehen. Das ist ein Datenschutzproblem, kein Sicherheitsproblem. Ein Angreifer, der eine PSBT liest, kann keine Mittel stehlen, weil die PSBT keine Private Keys enthält. Aber er erfährt, was der Nutzer ausgibt und wohin die Mittel gehen. Bei sensiblen Transactions sollte die PSBT über einen verschlüsselten Kanal oder auf physischen Medien transportiert werden, die der Nutzer kontrolliert.
