## La interfaz JSON-RPC

Bitcoin Core expone una interfaz JSON-RPC que permite a los programas consultar datos de la cadena de bloques y enviar transacciones. La herramienta de linea de comandos `bitcoin-cli` envia solicitudes a esta interfaz.

Cada dato en la cadena de bloques es consultable: cabeceras de bloques, bloques completos, transacciones individuales, saldos de direcciones, contenido del mempool, conexiones con pares y estadisticas de red.

Billeteras, exploradores de bloques, procesadores de pagos y herramientas de investigacion se comunican con Bitcoin Core a traves de esta API. La interfaz es el puente entre los datos crudos de la cadena de bloques y las aplicaciones que los presentan a los usuarios.
