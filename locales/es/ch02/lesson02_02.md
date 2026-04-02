## Propagación a través de la red

La billetera de Alice firma la transacción con su clave privada y la transmite a la red Bitcoin. Su billetera se conecta a varios nodos, que verifican la transacción de forma independiente: formato correcto, firma válida, entradas no gastadas, valores de salida que no exceden los valores de entrada.

Cada nodo que acepta la transacción la agrega a su mempool (una lista local de transacciones no confirmadas) y la retransmite a sus pares. En segundos, la transacción alcanza miles de nodos en la red. Ningún servidor central coordina esto. Cada nodo sigue las mismas reglas de forma independiente.

La billetera de Bob monitorea la red en busca de transacciones que paguen a su dirección. Cuando detecta la transacción de Alice, aparece como "no confirmada" en la billetera de Bob. El pago fue transmitido pero aún no se registró en un bloque.
