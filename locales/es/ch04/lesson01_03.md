## SHA-256 en Bitcoin

Bitcoin usa una funcion hash llamada SHA-256, disenada por la Agencia de Seguridad Nacional de Estados Unidos. SHA significa Algoritmo de Hash Seguro. 256 se refiere al tamano de salida: 256 bits, o 32 bytes.

SHA-256 aparece en todas partes en Bitcoin. Los identificadores de transaccion son hashes doble-SHA-256. La mineria implica encontrar entradas que produzcan hashes por debajo de un valor objetivo. La derivacion de direcciones usa SHA-256 combinado con RIPEMD-160 (una funcion hash diferente que produce una salida mas corta de 20 bytes). La combinacion de ambas, llamada HASH160, es lo que crea el hash de clave publica en una direccion.
