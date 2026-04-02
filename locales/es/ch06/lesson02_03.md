## Comisiones de transacción

Las transacciones de Bitcoin no tienen un campo de comisión explícito. La comisión es implícita: es la diferencia entre el valor total de todas las entradas y el valor total de todas las salidas.

Si tus entradas suman 100,000 satoshis y tus salidas suman 99,800 satoshis, la comisión es de 200 satoshis. Los mineros recogen esta comisión como incentivo por incluir tu transacción en un bloque.

Comisiones más altas significan confirmación más rápida. Cuando la red está congestionada, los mineros priorizan las transacciones con tasas de comisión más altas. Una transacción que paga muy poco puede esperar horas o días.

La tasa de comisión depende del peso de la transacción (medido en bytes virtuales, o vbytes), no de la cantidad enviada. SegWit introdujo unidades de peso donde los datos de testigo (firmas) tienen descuento comparados con otros datos de la transacción. Una transacción que envía 0.001 BTC puede costar la misma comisión que una que envía 1,000 BTC, si ambas tienen la misma estructura.
