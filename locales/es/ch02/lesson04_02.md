## El mempool

Cuando transmites una transacción, no entra en un bloque de inmediato. Ingresa al mempool: una sala de espera donde las transacciones sin confirmar permanecen hasta que un minero las incluye en un bloque.

Cada nodo mantiene su propio mempool. Las transacciones se propagan por la red en segundos, pero la confirmación (inclusión en un bloque) toma un promedio de 10 minutos. Durante períodos de alta actividad, el mempool crece y las transacciones con comisiones bajas pueden esperar más tiempo.

Una transacción en el mempool no está confirmada. Ha sido validada por los nodos (firmas correctas, entradas sin gastar) pero aún no se ha escrito en la blockchain.
