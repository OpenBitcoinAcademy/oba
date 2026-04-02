## P2PKH paso a paso

Esto es lo que sucede cuando un script P2PKH se ejecuta:

1. La firma y la clave pública se colocan en la pila (desde el script de desbloqueo).
2. OP_DUP duplica la clave pública en la pila.
3. OP_HASH160 aplica un hash al elemento superior (la clave pública duplicada) con SHA-256 y luego RIPEMD-160.
4. El hash de clave pública esperado se coloca en la pila (desde el script de bloqueo).
5. OP_EQUALVERIFY extrae los dos elementos superiores y verifica que sean iguales. Si difieren, el script falla.
6. OP_CHECKSIG extrae la firma y la clave pública, verifica la firma contra los datos de la transacción. Si es válida, coloca verdadero en la pila.

El resultado: verdadero en la pila. El gasto queda autorizado. La firma demuestra que el gastador controla la clave privada, sin revelarla.
