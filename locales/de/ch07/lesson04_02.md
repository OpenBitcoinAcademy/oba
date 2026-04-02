## MAST und Tapscript

Der Merkle Tree in Taproot ist ein Merklized Alternative Script Tree (MAST). Jedes Blatt des Baums ist ein anderes Script (eine andere Ausgabebedingung). Um über den Script Path auszugeben, enthüllt der Absender nur das Script, das er verwendet, und liefert einen Merkle Proof, dass es Teil des gebundenen Baums ist.

Nicht verwendete Scripts bleiben verborgen. Ein Taproot-Output mit 100 möglichen Ausgabebedingungen sieht auf der Blockchain gleich aus wie einer mit 1 Bedingung, solange der Script Path nicht verwendet wird. Unbenutzte Zweige bleiben privat.

Tapscript (BIP342) ist die Script-Sprache, die in Taproots Script-Path-Blättern verwendet wird. Sie ähnelt dem Legacy Script, hat aber Verbesserungen: OP_CHECKSIGADD ersetzt OP_CHECKMULTISIG (ermöglicht Batch-Validierung), Signature-Opcodes verwenden Schnorr statt ECDSA, und das Script-Größenlimit ist entfernt (ersetzt durch Gewichtslimits).
