## Pruebas Merkle

Una prueba Merkle permite verificar que una transaccion especifica esta incluida en un bloque sin descargar todas las transacciones. La prueba consiste en el hash de la transaccion, los hashes hermanos a lo largo del camino hacia la raiz, y la raiz Merkle del encabezado del bloque.

Para verificar, comienza con el hash de la transaccion. Aplica hash con su hermano. Aplica hash al resultado con el siguiente hermano. Continua hasta llegar a la raiz. Si tu raiz calculada coincide con la raiz Merkle en el encabezado del bloque, la transaccion esta en el bloque.

Un bloque con 4,000 transacciones requiere una prueba Merkle de solo unos 12 hashes (log2(4000)). Los clientes SPV usan pruebas Merkle para verificar transacciones sin descargar bloques completos.
