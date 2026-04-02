## Locking Scripts

Jeder Transaction-Output enthält ein Locking Script, genannt scriptPubKey. Dieses Script definiert die Bedingungen, die erfüllt sein müssen, um den Output auszugeben.

Das am leichtesten verständliche Locking Script ist Pay-to-Public-Key-Hash (P2PKH). Es besagt: "Um diesen Output auszugeben, beweise, dass du den Private Key besitzt, der zu diesem Public Key Hash gehört." Moderne Transactions verwenden neuere Formate (P2WPKH für SegWit, P2TR für Taproot), aber P2PKH zeigt das Kernkonzept am deutlichsten.

In Script-Notation sieht ein P2PKH-Lock so aus: OP_DUP OP_HASH160 <pubkey_hash> OP_EQUALVERIFY OP_CHECKSIG. Jedes Teil ist ein Befehl, den die Bitcoin-Script-Engine ausführt.
