## Unidades de peso

Segwit introdujo unidades de peso para reemplazar la métrica de tamaño basada en bytes. Cada byte de datos no-witness cuesta 4 unidades de peso. Cada byte de datos witness (firmas) cuesta 1 unidad de peso.

Este descuento incentiva el uso de transacciones segwit, que almacenan las firmas en la estructura witness. Una transacción segwit usa menos peso que una transacción legacy con la misma cantidad de entradas y salidas.

Los bytes virtuales (vbytes) convierten el peso a un equivalente en bytes: vbytes = peso / 4. Una transacción con 600 unidades de peso tiene 150 vbytes. Las tasas de comisión expresadas en sat/vB incorporan el descuento segwit automáticamente.
