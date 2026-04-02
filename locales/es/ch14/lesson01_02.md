## Por que funciona el ajuste

El ajuste vincula el arbol de scripts a la clave publica criptograficamente. Cambiar cualquier script en el arbol cambia la raiz de Merkle, lo que cambia el ajuste, lo que cambia Q. Un script comprometido en Q al momento de crear la salida no puede modificarse despues.

La clave privada se ajusta de forma identica: tweaked_privkey = (privkey + t) mod n. El titular de la clave privada interna puede calcular la clave privada ajustada y firmar directamente. Este es el gasto por ruta de clave.

Los tagged hashes previenen colisiones entre diferentes usos de la funcion hash. La etiqueta "TapTweak" tiene separacion de dominio: SHA256(SHA256("TapTweak") || SHA256("TapTweak") || data). Etiquetas diferentes producen salidas hash diferentes para los mismos datos, eliminando ataques entre protocolos.
