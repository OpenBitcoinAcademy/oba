## Transacciones en conflicto

Alice puede crear dos transacciones que gasten la misma salida: una pagando a Bob y otra pagando a Carol. Ambas son firmas válidas, pero solo una puede incluirse en la cadena de bloques (la salida solo puede gastarse una vez).

Los mineros deciden qué transacción en conflicto incluir. Por defecto, la mayoría de los mineros siguen una política de "primero visto" e incluyen la transacción que recibieron primero. Pero esto es una política, no una regla de consenso. Un minero puede incluir cualquier transacción válida.

Por esto Bob debe esperar confirmaciones antes de considerar un pago como final. La comisión que Alice paga incentiva a los mineros a incluir su transacción rápidamente, reduciendo la ventana para transacciones en conflicto.
