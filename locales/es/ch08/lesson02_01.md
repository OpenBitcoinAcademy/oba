## El algoritmo ECDSA

ECDSA (Elliptic Curve Digital Signature Algorithm) fue el esquema de firma original de Bitcoin. Opera sobre la curva elíptica secp256k1, la misma curva usada para la generación de claves.

Para firmar un mensaje, el algoritmo toma la clave privada y un hash del mensaje, los combina con un número aleatorio (llamado k, o el nonce), y produce dos valores: r y s. Juntos, (r, s) forman la firma.

Para verificar, el algoritmo toma la clave pública, el hash del mensaje y la firma (r, s). Realiza cálculos de curva elíptica para comprobar si la firma es consistente con la clave pública y el mensaje. No se necesita la clave privada para la verificación.
