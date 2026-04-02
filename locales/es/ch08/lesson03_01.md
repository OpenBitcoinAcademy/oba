## Schnorr: una firma más sencilla

Las firmas Schnorr (BIP340) se activaron con la actualización Taproot en noviembre de 2021. Se usan para transacciones segwit v1 (P2TR).

El algoritmo Schnorr es matemáticamente más sencillo que ECDSA. Produce una firma fija de 64 bytes (comparado con los 71-73 bytes variables de ECDSA). La ecuación de verificación es lineal, lo que permite la agregación de firmas: múltiples firmas pueden combinarse en una sola firma del mismo tamaño.

Las firmas Schnorr tienen una reducción de seguridad demostrable al problema del logaritmo discreto, una garantía de seguridad más fuerte que ECDSA.
