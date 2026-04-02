## Firmas en transacciones de Bitcoin

En Bitcoin, las firmas aparecen en las entradas de la transacción (para transacciones legacy) o en la estructura witness (para transacciones segwit). Cada entrada que gasta un UTXO requiere una firma que demuestre que el gastador controla la clave que bloqueó esa salida.

Múltiples partes pueden colaborar en una sola transacción. Cada parte firma solo la entrada que controla. Las firmas son independientes: un firmante no necesita acceso a la clave privada de otro.

Bitcoin usa dos algoritmos de firma: ECDSA (Elliptic Curve Digital Signature Algorithm) para transacciones legacy y segwit v0, y firmas Schnorr para transacciones segwit v1 (Taproot).
