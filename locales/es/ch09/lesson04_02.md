## Transaction pinning

Transaction pinning es un ataque donde una parte maliciosa crea una transacción hija con comisión baja que hace costoso o imposible para la parte honesta incrementar la comisión de la madre.

En un protocolo de dos partes (como un canal de Lightning Network), una parte podría transmitir un descendiente grande y con comisión baja de una transacción compartida. El intento de CPFP de la otra parte tendría que pagar por el descendiente grande del atacante, haciendo el incremento de comisión prohibitivamente caro.

Los anchor outputs son una contramedida. Cada parte recibe una salida pequeña en la transacción compartida que puede gastar con CPFP. Las reglas limitan cuántos descendientes puede tener cada anchor output, previniendo el ataque de pinning.
