## Scripts de bloqueo

Cada salida de transacción contiene un script de bloqueo, llamado scriptPubKey. Este script define las condiciones que deben cumplirse para gastar la salida.

El script de bloqueo más sencillo de entender es Pay-to-Public-Key-Hash (P2PKH). Dice: "Para gastar esta salida, demuestra que posees la clave privada correspondiente a este hash de clave pública." Las transacciones modernas usan formatos más nuevos (P2WPKH para SegWit, P2TR para Taproot), pero P2PKH demuestra el concepto central con mayor claridad.

En notación de script, un bloqueo P2PKH se ve así: OP_DUP OP_HASH160 <pubkey_hash> OP_EQUALVERIFY OP_CHECKSIG. Cada pieza es una instrucción que el motor de scripts de Bitcoin ejecuta.
