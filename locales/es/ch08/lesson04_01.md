## Tipos de SIGHASH

Una firma no se compromete con la transacción completa por defecto. El flag SIGHASH añadido a cada firma especifica qué partes de la transacción cubre la firma.

**SIGHASH_ALL** (0x01) es el valor por defecto. La firma se compromete con todas las entradas y todas las salidas. Cambiar cualquier parte de la transacción invalida la firma.

**SIGHASH_NONE** (0x02) se compromete con todas las entradas pero ninguna salida. Cualquiera puede cambiar las salidas después de firmar. Se usa en protocolos colaborativos poco comunes.

**SIGHASH_SINGLE** (0x03) se compromete con todas las entradas pero solo la salida en el mismo índice que la entrada siendo firmada. Las demás salidas pueden cambiarse.

El modificador ANYONECANPAY puede combinarse con cualquiera de estos, permitiendo firmas que se comprometen con una sola entrada.
