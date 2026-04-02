## Taproot-Felder und Versionsunterschiede

BIP 371 fügte Taproot-spezifische Felder zum PSBT-Format hinzu. TAP_KEY_SIG (Input Key-Typ 0x13) speichert eine Schnorr Signature für einen Key Path Spend. TAP_SCRIPT_SIG (Input Key-Typ 0x14) speichert eine Schnorr Signature für ein bestimmtes Script Leaf, geschlüsselt nach X-only Public Key und Leaf Hash.

TAP_LEAF_SCRIPT (Input Key-Typ 0x15) liefert das Script, seine Leaf-Version und den Control Block, der für Script Path Spending benötigt wird. TAP_BIP32_DERIVATION (Input Key-Typ 0x16) erweitert das Standard-BIP-32-Ableitungsfeld um eine Liste von Leaf Hashes, für die der Key signieren kann. TAP_INTERNAL_KEY (Input Key-Typ 0x17) zeichnet den internen Public Key vor dem Tweaking auf.

Auf der Output-Seite lassen TAP_INTERNAL_KEY (Output Key-Typ 0x05) und TAP_BIP32_DERIVATION (Output Key-Typ 0x07) Signer verifizieren, dass Taproot Change Outputs zum selben Wallet gehören. Der Signer kann den getweakten Key aus dem Internal Key neu ableiten und bestätigen, dass er mit dem scriptPubKey des Outputs übereinstimmt.

PSBTv2 (BIP 370) unterscheidet sich von v0 in der Struktur, nicht im Konzept. In v0 liegt die unsignierte Transaction als einzelner serialisierter Blob in der globalen Map. In v2 werden Inputs und Outputs durch Per-Map-Felder beschrieben: PREVIOUS_TXID, OUTPUT_INDEX, SEQUENCE für Inputs; AMOUNT und SCRIPT für Outputs. Constructors können Inputs und Outputs inkrementell hinzufügen, ohne jedes Mal die gesamte Transaction neu zu serialisieren.
