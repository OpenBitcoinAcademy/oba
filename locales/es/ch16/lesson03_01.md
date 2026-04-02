## Los tres mapas

Un PSBT contiene tres tipos de mapas clave-valor: un mapa global, un mapa por entrada y un mapa por salida. Cada mapa almacena pares clave-valor tipados terminados por un byte 0x00. El tipo de clave determina el significado de la entrada. El payload de clave y el payload de valor llevan los datos.

El mapa global contiene datos que aplican a toda la transaccion. En PSBTv0, la entrada global mas importante es UNSIGNED_TX (tipo de clave 0x00), que contiene la transaccion completa sin firmar. En PSBTv2, el mapa global contiene TX_VERSION, FALLBACK_LOCKTIME, INPUT_COUNT y OUTPUT_COUNT como campos separados. La transaccion sin firmar se reconstruye a partir de mapas por-entrada y por-salida en lugar de almacenarse completa.

Las entradas globales tambien incluyen XPUB (tipo de clave 0x01), que mapea una clave publica extendida a su ruta de derivacion BIP 32. Esto permite a un firmante verificar que las salidas de cambio derivan de la misma raiz del monedero sin necesidad de acceder al descriptor completo del monedero.
