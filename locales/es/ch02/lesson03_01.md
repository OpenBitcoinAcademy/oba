## La lotería de la minería

La minería es una lotería descentralizada. Cada minero crea un bloque candidato a partir de las transacciones en su mempool. Hashean el encabezado del bloque repetidamente, cambiando un número llamado nonce cada vez, buscando un valor de hash por debajo de un objetivo establecido por la red.

Encontrar un hash válido requiere billones de intentos. Verificar un hash válido requiere calcularlo una sola vez. Esta asimetría es el núcleo del proof of work: crear un bloque válido es costoso, verificar uno es barato.

El minero que encuentra un hash válido primero transmite el bloque. Los demás mineros lo verifican, lo aceptan e inmediatamente comienzan a trabajar en el siguiente bloque. El ganador cobra la recompensa del bloque: bitcoins recién creados (el subsidio) más la suma de todas las comisiones de transacción en el bloque.
