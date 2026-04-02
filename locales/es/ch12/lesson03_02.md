## Reglas de validacion de bloques

Cuando un nodo recibe un nuevo bloque, verifica docenas de criterios antes de aceptarlo. El hash de la cabecera del bloque debe estar por debajo del objetivo actual. La marca de tiempo debe estar dentro de limites aceptables. La primera transaccion debe ser una coinbase, y la salida de la coinbase no debe exceder el subsidio permitido mas las comisiones. Cada otra transaccion debe ser valida segun las reglas de script.

La raiz de Merkle del bloque debe coincidir con el arbol de hashes de todas las transacciones incluidas. El tamano del bloque no debe exceder el limite de consenso. El bloque debe referenciar un bloque padre valido que el nodo ya tenga.

Si alguna verificacion falla, el nodo rechaza el bloque y puede desconectarse del par que lo envio. No hay proceso de apelacion. Un bloque es valido o no lo es. Esta validacion estricta es lo que impide que los mineros creen bitcoin de la nada o gasten monedas que no poseen.
