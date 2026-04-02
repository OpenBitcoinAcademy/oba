## Requisitos de recursos

Ejecutar un nodo completo requiere espacio en disco, ancho de banda y tiempo. La descarga inicial de la cadena de bloques supera los 500 GB a fecha de 2024 y crece aproximadamente 150 MB por dia.

Bitcoin Core puede ejecutarse en modo podado, descartando datos de bloques antiguos despues de la verificacion y conservando solo el conjunto UTXO y los bloques recientes. Esto reduce el almacenamiento a menos de 10 GB. El nodo aun verifica toda la historia durante la sincronizacion inicial.

Un nodo completo puede funcionar en hardware tan modesto como una Raspberry Pi. La sincronizacion inicial tarda mas en hardware lento, pero una vez sincronizado, los datos diarios son manejables con la mayoria de las conexiones de internet domesticas.
