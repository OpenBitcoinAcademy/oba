## Construyendo la transacción

Una transacción legacy tiene cuatro campos: versión, entradas, salidas y locktime. Las transacciones segwit modernas agregan tres más: un marcador, una bandera y una estructura de testigo que contiene los datos de autorización (firmas) separados de las entradas.

**Versión** es un número (actualmente 1 o 2) que indica a los nodos qué reglas de validación aplicar.

**Entradas** listan los UTXOs que se van a gastar. Cada entrada especifica el ID de la transacción anterior, el índice de la salida dentro de esa transacción, un script de entrada y un número de secuencia.

**Salidas** listan los nuevos UTXOs que se crean. Cada salida especifica una cantidad en satoshis y un script de bloqueo.

**Locktime** generalmente es cero. Cuando se configura con una altura de bloque futura o una marca de tiempo, la transacción no puede incluirse en un bloque hasta ese punto.

La transacción se serializa en una secuencia de bytes, se aplica doble hash con SHA-256, y el resultado es el ID de la transacción.
