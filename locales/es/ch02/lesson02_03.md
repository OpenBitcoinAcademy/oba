## Del mempool a la blockchain

La transacción de Alice permanece en el mempool, esperando que un minero la incluya en un bloque. Los mineros seleccionan transacciones de su mempool, priorizando aquellas con tasas de comisión más altas por unidad de peso.

Unos cinco minutos después, un minero encuentra una solución válida de proof of work para un nuevo bloque que incluye la transacción de Alice. El minero transmite el bloque. Cada nodo completo verifica el bloque de forma independiente: proof of work válido, transacciones válidas, formato correcto. Los nodos que aceptan el bloque lo agregan a su copia de la blockchain.

La transacción de Alice ahora tiene una confirmación. La billetera de Bob actualiza el estado. Cada bloque subsiguiente construido sobre este agrega otra confirmación. Después de seis confirmaciones (aproximadamente una hora), la transacción se considera altamente segura contra reversión.
