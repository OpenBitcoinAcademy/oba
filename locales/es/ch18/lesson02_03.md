## Decrementos de timelock

Cada salto en la cadena tiene un timelock mas corto que el salto anterior. Alice da a Bob 40 bloques. Bob da a Carol 30 bloques. Este decremento (llamado delta CLTV) asegura que cada nodo de reenvio tenga tiempo para reclamar fondos de su salto siguiente antes de que expire su HTLC anterior.

Si Carol no revela R dentro de 30 bloques, el HTLC de Bob expira y los fondos regresan a Bob. Bob todavia tiene 10 bloques para liquidar con Alice. Si Bob no revela R a Alice dentro de 40 bloques, los fondos de Alice regresan a ella.

El decremento previene un ataque de temporalizacion donde un nodo aguas abajo retrasa la revelacion de la preimagen hasta que el HTLC aguas arriba expira, atrapando al nodo de reenvio.
