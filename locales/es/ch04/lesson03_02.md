## Los pasos de derivacion

Crear una direccion Bitcoin a partir de una clave publica requiere varios pasos de hashing.

Primero, haz hash de la clave publica con SHA-256. Luego haz hash de ese resultado con RIPEMD-160. Esto produce un hash de 20 bytes llamado hash de clave publica, o HASH160.

A continuacion, agrega un byte de version al frente. Este byte identifica la red (mainnet vs testnet) y el tipo de direccion.

Finalmente, codifica el resultado. Las direcciones legacy usan codificacion Base58Check, que agrega una suma de verificacion de 4 bytes y codifica en un alfabeto legible que evita caracteres confusos como 0/O y l/1. Las direcciones SegWit mas nuevas usan codificacion Bech32, que es solo en minusculas y tiene mejor deteccion de errores.
