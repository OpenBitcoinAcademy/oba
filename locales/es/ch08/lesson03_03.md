## Schnorr vs ECDSA

Ambos algoritmos usan la curva secp256k1. Ambos proporcionan seguridad equivalente contra recuperación de clave. Las diferencias son prácticas.

Las firmas Schnorr son más pequeñas (64 bytes vs 71-73). Schnorr soporta agregación nativa de firmas. Schnorr tiene una ecuación de verificación más sencilla. Schnorr tiene una prueba de seguridad formal.

ECDSA fue elegido para Bitcoin en 2008 porque Schnorr estaba patentado hasta 2008, y el estado de la patente era incierto cuando Satoshi diseñó el sistema. Con la patente expirada y Taproot activado, las transacciones nuevas pueden usar Schnorr. ECDSA sigue soportado por compatibilidad.
