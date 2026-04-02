## Wrapped SegWit (P2SH-P2WPKH)

Nicht alle Wallets unterstützten native Segwit-Addresses (bc1...), als Segwit aktiviert wurde. Wrapped Segwit bietet Rückwärtskompatibilität, indem es das Segwit-Programm in ein P2SH-Redeem-Script platziert.

Der Output sieht aus wie eine Standard-P2SH-Address (beginnend mit "3"). Das Redeem Script enthält: OP_0 <20-byte-pubkey-hash>. Beim Ausgeben enthüllt das Input Script das Redeem Script, und der Witness liefert die Signature und den Public Key.

Wrapped Segwit ist ein Übergangsformat. Native Segwit (bc1q...-Addresses) wird für neue Transactions bevorzugt, weil es kleiner, günstiger und dank Bech32-Kodierung besser bei der Fehlererkennung ist.
