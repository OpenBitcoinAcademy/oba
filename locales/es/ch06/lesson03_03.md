## Varints y endianness

Bitcoin usa dos convenciones de codificación en toda su serialización.

Los **varints** (enteros de longitud variable) codifican conteos y longitudes. Un valor menor a 253 cabe en un byte. Los valores de 253 a 65535 usan un prefijo 0xFD seguido de 2 bytes. Los valores hasta ~4 mil millones usan 0xFE seguido de 4 bytes. Los valores mayores usan 0xFF seguido de 8 bytes.

El orden de bytes **little-endian** coloca el byte menos significativo primero. La versión 1 se serializa como 01 00 00 00, no como 00 00 00 01. Los valores en satoshis, las alturas de bloque y la mayoría de los campos numéricos usan little-endian.

Los hashes de transacción (txids) se muestran en orden de bytes invertido por convención. El hash doble SHA-256 produce bytes en un orden, pero los exploradores de bloques y Bitcoin Core los muestran invertidos. Esta es una convención de visualización, no una regla de serialización.
