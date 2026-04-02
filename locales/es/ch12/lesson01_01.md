## La mineria como loteria descentralizada

Bitcoin no tiene una autoridad central que decida que transacciones van en los bloques. En su lugar, los mineros compiten en una loteria computacional. El ganador obtiene el derecho de agregar el siguiente bloque a la cadena.

Cada minero toma un conjunto de transacciones no confirmadas, las ensambla en un bloque candidato y aplica un hash a la cabecera del bloque usando SHA-256 (aplicado dos veces). El resultado debe ser un numero menor que el objetivo actual. La mayoria de los intentos fallan. El minero cambia el campo nonce en la cabecera y aplica el hash de nuevo.

Este proceso se repite miles de millones de veces por segundo entre todos los mineros del mundo. El minero que encuentra un hash valido primero transmite el bloque. Otros nodos lo verifican de forma independiente. Nadie necesita permiso para minar, y nadie puede predecir quien ganara la siguiente vez.
