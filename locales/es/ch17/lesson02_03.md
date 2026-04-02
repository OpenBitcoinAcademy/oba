## De porciones a una firma estandar

El coordinador recopila t porciones de firma y las suma. El resultado es una unica firma Schnorr: un par (R, s) de 64 bytes, donde R es el punto de nonce agregado y s es la suma de las porciones de firma. Esta firma es valida bajo la clave publica agregada del grupo.

Un nodo completo de Bitcoin que verifica esta transaccion ejecuta la ecuacion estandar de verificacion Schnorr de BIP 340. Verifica una firma contra una clave publica. El nodo no tiene forma de saber si la firma fue producida por tres firmantes, cinco firmantes o un solo firmante. La verificacion es identica.

Por esto FROST es poderoso para Bitcoin: no requiere cambios de consenso. Taproot y BIP 340 ya aceptan las firmas que FROST produce. La complejidad de la firma umbral vive completamente en el software de los firmantes. La blockchain y todos los verificadores permanecen sin enterarse.
