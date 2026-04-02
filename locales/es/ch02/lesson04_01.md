## Qué es un bloque?

Un bloque es un lote de transacciones agrupadas y añadidas a la blockchain. Cada bloque contiene un encabezado y una lista de transacciones.

El encabezado del bloque tiene seis campos: versión, hash del bloque anterior, raíz de Merkle (un hash de todas las transacciones del bloque), marca de tiempo, objetivo de dificultad y nonce. El hash del bloque anterior enlaza cada bloque con el anterior, formando una cadena.

Los mineros compiten por encontrar un valor nonce que haga que el hash del encabezado del bloque quede por debajo del objetivo de dificultad. Esto requiere billones de intentos. El primer minero que encuentra un nonce válido difunde el bloque. Los demás nodos lo verifican y lo añaden a su copia de la cadena.
