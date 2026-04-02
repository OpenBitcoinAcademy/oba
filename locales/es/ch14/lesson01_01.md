## La clave interna y el ajuste

Una salida Taproot bloquea fondos a una clave publica ajustada Q. Esta clave se deriva de dos entradas: una clave publica interna P y una raiz de Merkle opcional de un arbol de scripts.

El valor de ajuste t se calcula como: t = tagged_hash("TapTweak", P || merkle_root). La clave ajustada es: Q = P + t * G, donde G es el punto generador en secp256k1.

Si no hay arbol de scripts, el ajuste usa solo P: t = tagged_hash("TapTweak", P). La salida aun se compromete con la clave interna, pero no tiene scripts incrustados.

En la cadena, el scriptPubKey es: OP_1 seguido de la coordenada X de 32 bytes de Q. Son 34 bytes, el mismo tamano que una salida P2WSH. Ningun observador puede distinguir si Q contiene un arbol de scripts incrustado o no.
