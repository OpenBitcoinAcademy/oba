## El requisito de aleatoriedad

Tanto ECDSA como Schnorr requieren un número aleatorio (nonce) durante la creación de la firma. La seguridad de todo el sistema depende de esta aleatoriedad.

Si el mismo nonce se usa para firmar dos mensajes diferentes con la misma clave privada, la clave privada puede calcularse a partir de las dos firmas. Este no es un riesgo teórico. En 2013, un fallo en el generador de números aleatorios de Android causó que múltiples billeteras de Bitcoin reutilizaran nonces, y se robaron fondos.

Las implementaciones modernas usan nonces determinísticos (RFC 6979 para ECDSA, BIP340 para Schnorr), derivando el nonce de la clave privada y el hash del mensaje. Esto elimina la dependencia de un generador de números aleatorios y previene la reutilización de nonces.
