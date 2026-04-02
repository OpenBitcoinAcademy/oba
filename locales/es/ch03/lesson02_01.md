## Que hace un nodo completo

Un nodo completo descarga y verifica de forma independiente cada bloque en la cadena de bloques, comenzando desde el bloque genesis en 2009. Comprueba cada transaccion contra cada regla de consenso. Mantiene su propia copia del conjunto UTXO (todas las salidas de transacciones no gastadas).

Cuando ejecutas un nodo completo, no confias en nadie. Tu nodo verifica toda la historia de la red por si mismo. Si un bloque o transaccion viola alguna regla, tu nodo lo rechaza, sin importar cuantos otros nodos lo acepten.

Un nodo completo tambien retransmite transacciones y bloques validos a otros nodos, contribuyendo a la salud y resiliencia de la red.
