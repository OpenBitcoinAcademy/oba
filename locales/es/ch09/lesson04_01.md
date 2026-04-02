## Package relay

Históricamente, Bitcoin Core evaluaba cada transacción de forma independiente para la admisión al mempool. Una transacción madre con comisión baja era rechazada aunque una transacción hija con comisión alta dependiera de ella.

Package relay cambia esto. Los nodos evalúan grupos de transacciones relacionadas en conjunto, aceptando una madre con comisión baja si su hija eleva la tasa de comisión combinada por encima del umbral.

Esto mejora la fiabilidad de CPFP. Sin package relay, la madre podría no propagarse a los mineros, haciendo inútil a la hija. Con package relay, madre e hija viajan juntas a través de la red.
