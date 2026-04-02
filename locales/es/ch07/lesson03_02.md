## SegWit Envuelto (P2SH-P2WPKH)

No todas las billeteras soportaban direcciones segwit nativas (bc1...) cuando segwit se activo por primera vez. SegWit envuelto proporciona compatibilidad hacia atras colocando el programa segwit dentro de un redeem script P2SH.

La salida luce como una direccion P2SH estandar (comenzando con "3"). El redeem script contiene: OP_0 <hash-de-clave-publica-de-20-bytes>. Al gastar, el script de entrada revela el redeem script, y el testigo proporciona la firma y la clave publica.

SegWit envuelto es un formato de transicion. SegWit nativo (direcciones bc1q...) es preferido para nuevas transacciones porque es mas pequeno, mas barato, y tiene mejor deteccion de errores mediante la codificacion bech32.
