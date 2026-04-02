## Das Gesamtbild

Threshold Signatures verändern die Funktionsweise der Bitcoin-Verwahrung. Der alte Ansatz legte Spending Policies on-chain offen: Jeder konnte die Keys zählen, den Threshold sehen und Multisig-Wallets über Transactions hinweg verfolgen. FROST und ChillDKG verlagern all diese Komplexität off-chain.

Die Gruppe verhandelt ihren Threshold und erzeugt Keys mit ChillDKG. Beliebige t Unterzeichner erzeugen eine Signature mit dem FROST-Protokoll. Die Blockchain sieht einen Taproot-Key-Path-Spend. Verifizierer prüfen eine Signature gegen einen Key. Die On-Chain-Kosten sind konstant, unabhängig davon, wie viele Teilnehmer beteiligt sind: 64 Bytes für die Signature, 32 Bytes für den Key.

Das ist möglich, weil Taproot und BIP-340-Schnorr-Verifikation bereits im Bitcoin-Netzwerk aktiv sind. Kein Soft Fork erforderlich. Keine neuen Opcodes. Das Threshold-Signing-Protokoll läuft vollständig in den Wallets der Teilnehmer. Die Konsensschicht verifiziert das Ergebnis mit denselben Regeln, die sie für jeden anderen Taproot-Spend verwendet.
