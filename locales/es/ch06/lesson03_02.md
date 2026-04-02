## Byte por byte

**Versión** (4 bytes): 01000000 significa versión 1. 02000000 significa versión 2 (habilita las restricciones de secuencia de BIP68). Codificación little-endian: el byte menos significativo va primero.

**Entradas**: un conteo varint seguido de cada entrada. Cada entrada contiene el hash de la transacción previa (32 bytes, invertido), el índice de salida (4 bytes), la longitud del script de entrada (varint), el script de entrada (variable) y el número de secuencia (4 bytes).

**Salidas**: un conteo varint seguido de cada salida. Cada salida contiene el valor en satoshis (8 bytes, little-endian) y la longitud del script de salida (varint) seguida del script de salida.

**Testigo** (solo segwit): para cada entrada, un conteo de elementos de testigo seguido de la longitud y los datos de cada elemento. Las entradas legacy tienen cero elementos de testigo.

**Locktime** (4 bytes): normalmente 00000000. Un valor distinto de cero restringe cuándo se puede minar la transacción.

El identificador de transacción (txid) es el hash doble SHA-256 de la transacción serializada, con los datos de testigo excluidos (para segwit) o incluidos (para legacy).
