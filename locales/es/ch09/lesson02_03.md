## Estimación de comisiones

El software de billetera estima las tasas de comisión analizando el mempool: el conjunto de transacciones no confirmadas esperando ser incluidas en un bloque.

Un enfoque común: observar las tasas de comisión de las transacciones incluidas en bloques recientes y las tasas de las transacciones que aún esperan. Establecer una tasa de comisión suficiente para competir con la transacción de menor tasa en el bloque más reciente, con un margen de seguridad.

La estimación de comisiones es imperfecta. La demanda fluctúa. Un pico de uso puede hacer que una tasa de comisión previamente adecuada resulte insuficiente. Algunas billeteras permiten a los usuarios establecer tasas personalizadas para mayor control.
