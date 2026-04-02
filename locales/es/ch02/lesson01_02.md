## Las entradas referencian salidas anteriores

Cada entrada apunta a una salida específica de una transacción anterior. Lo hace con dos datos: el ID de la transacción (un hash) y el índice de la salida (cuál salida dentro de esa transacción).

Cuando gastas bitcoin, demuestras que controlas la clave capaz de desbloquear una salida anterior. En transacciones legacy, la prueba (firma y clave pública) vive en el script de entrada. En transacciones segwit modernas, la prueba vive en una estructura de testigo separada, y el script de entrada está vacío.

Una vez que una salida es referenciada por una entrada, queda gastada. No se puede gastar de nuevo. Así es como Bitcoin previene el doble gasto sin una autoridad central.
