## Objetivo, dificultad y el nonce

El objetivo es un numero de 256 bits. Un hash de bloque valido debe ser numericamente menor que el objetivo. Un objetivo mas bajo significa que existen menos hashes validos, lo que hace mas dificil encontrar uno. La dificultad es la inversa del objetivo: cuando el objetivo baja, la dificultad sube.

La cabecera del bloque contiene un campo nonce de 32 bits. Los mineros incrementan este valor con cada intento de hash, recorriendo todas las $2^{32}$ (aproximadamente 4,300 millones de) posibilidades. Los mineros modernos agotan el espacio del nonce en menos de un segundo.

Cuando el espacio del nonce se agota, los mineros modifican otros campos para crear nuevas entradas de hash. La tecnica mas comun cambia el campo de nonce extra de la transaccion coinbase. Esto cambia la raiz de Merkle en la cabecera del bloque, dando al minero un nuevo conjunto de $2^{32}$ nonces para probar.
