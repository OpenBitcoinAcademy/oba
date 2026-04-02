## El formato de transmisión

Una transacción de Bitcoin se serializa en una secuencia de bytes para su transmisión a través de la red y su almacenamiento en bloques. El formato de serialización define el orden exacto de bytes y la codificación de cada campo.

Una transacción legacy tiene cuatro campos de nivel superior serializados en orden: versión (4 bytes, little-endian), entradas (variable), salidas (variable) y locktime (4 bytes, little-endian).

Una transacción segwit agrega tres campos: después de la versión, un byte marker (0x00) y un byte flag (0x01) señalan que siguen datos de testigo. La estructura de testigo aparece después de las salidas y antes del locktime. Un parser que ve un byte cero donde debería estar la cantidad de entradas sabe que está leyendo el formato extendido (segwit).
