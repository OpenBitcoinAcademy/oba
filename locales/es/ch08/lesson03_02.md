## Agregación de multifirma

Con ECDSA, una transacción que requiere tres firmantes necesita tres firmas separadas, cada una de 71-73 bytes. Las firmas se verifican individualmente.

Con Schnorr, los mismos tres firmantes pueden combinar sus firmas en una sola firma agregada de 64 bytes usando protocolos como MuSig2. En la blockchain, la firma agregada luce idéntica a una firma de un solo firmante. Ningún observador puede saber cuántas partes participaron.

Esto tiene dos beneficios. Las transacciones con múltiples firmantes usan menos espacio de bloque (comisiones más bajas). Y las transacciones multifirma se vuelven indistinguibles de las transacciones de firma única, mejorando la privacidad.
