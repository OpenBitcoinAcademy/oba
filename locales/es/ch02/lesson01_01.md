## Cómo se ve una transacción

Una transacción de Bitcoin es una estructura de datos que transfiere bitcoin de un propietario a otro. Contiene entradas (que referencian salidas anteriores para gastar), salidas (que crean nuevas cantidades gastables) y datos de autorización (firmas que prueban que el gastador controla las claves).

Las entradas indican de dónde proviene el bitcoin. Cada entrada hace referencia a una salida de una transacción anterior que aún no se ha gastado. Esta salida sin gastar se llama UTXO (unspent transaction output).

Las salidas indican a dónde va el bitcoin. Cada salida especifica una cantidad en satoshis y una condición de bloqueo (un script) que determina quién puede gastarla.

Una transacción consume salidas anteriores y crea nuevas. No hay nada "almacenado en una cuenta". Bitcoin rastrea la propiedad a través de una cadena de salidas, cada una bloqueada a una clave específica.
