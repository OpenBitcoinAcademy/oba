## Replace By Fee (RBF)

Si una transacción está atascada en el mempool porque su tasa de comisión es demasiado baja, el emisor puede reemplazarla con una nueva versión que pague una comisión más alta. Esto es Replace By Fee (RBF), definido en BIP125.

Para usar RBF, la transacción original debe señalar que es reemplazable estableciendo el número de secuencia de al menos una entrada a un valor menor que 0xFFFFFFFE. El software de billetera maneja esto automáticamente.

La transacción de reemplazo debe pagar una comisión total más alta que la original. Puede cambiar las salidas (pagando un monto diferente o agregando/eliminando destinatarios) siempre que gaste al menos una de las mismas entradas.
