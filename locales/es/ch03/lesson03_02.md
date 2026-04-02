## Explorando la cadena de bloques

Con `bitcoin-cli`, puedes recuperar cualquier bloque por su altura o hash. El comando `getblock` devuelve los campos de la cabecera del bloque, la lista de IDs de transacciones y metadatos como el numero de confirmaciones.

El comando `getrawtransaction` devuelve los datos hexadecimales crudos de una transaccion o su representacion JSON decodificada, mostrando entradas, salidas, scripts y datos de testigo.

El comando `getblockchaininfo` reporta la altura actual de la cadena, la dificultad, el progreso de verificacion y si el nodo aun esta sincronizando. Estos comandos convierten conceptos abstractos de la cadena de bloques en datos concretos e inspeccionables.
