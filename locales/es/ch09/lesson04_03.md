## Fee sniping

Fee sniping es un ataque teórico donde un minero, en lugar de construir sobre el último bloque, re-mina un bloque anterior para robar sus comisiones de transacción. Si las comisiones en el bloque N son suficientemente altas, un minero podría encontrar más rentable bifurcar la cadena en el bloque N-1 y recaudar esas comisiones.

Bitcoin Core se defiende contra esto estableciendo el locktime de las nuevas transacciones a la altura del bloque actual. Una transacción bloqueada al bloque 800,000 no puede incluirse en el bloque 799,999. Esto hace menos rentable re-minar bloques antiguos, porque las nuevas transacciones creadas desde entonces no estarían disponibles.

Fee sniping no ha ocurrido en la red Bitcoin. La defensa es precautoria.
