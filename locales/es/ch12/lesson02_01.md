## La transaccion coinbase

Cada bloque comienza con una transaccion especial llamada transaccion coinbase. No tiene entradas de transacciones previas. En su lugar, crea bitcoin nuevo como recompensa para el minero que encontro el bloque.

La transaccion coinbase tiene una entrada con una referencia nula (ninguna salida previa siendo gastada). El minero llena el campo de script de esta entrada con datos arbitrarios, incluyendo el nonce extra usado durante la mineria. Satoshi Nakamoto incrusto un famoso titular de periodico en la coinbase del bloque genesis: "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks."

Las salidas de la transaccion coinbase pagan al minero. El valor total de la salida no debe exceder el subsidio de bloque mas la suma de todas las comisiones de transaccion del bloque. Cualquier valor no reclamado se destruye.
