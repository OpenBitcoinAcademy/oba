## Filtros de bloques compactos

**BIP 157** y **BIP 158** definen un mejor enfoque llamado Compact Block Filters. En lugar de que el cliente envie su filtro al servidor, el servidor construye un filtro para cada bloque y lo envia al cliente.

Cada filtro es una representacion compacta de todas las direcciones y scripts en un bloque. El cliente descarga el filtro y verifica localmente si alguna de sus direcciones aparece. Si encuentra una coincidencia, el cliente descarga el bloque completo para extraer las transacciones relevantes.

La ventaja de privacidad es que el servidor nunca conoce las direcciones que le interesan al cliente. El servidor envia el mismo filtro a todos los clientes. El cliente hace todo el emparejamiento localmente.

El costo de ancho de banda es mayor que con los filtros Bloom porque el cliente descarga un filtro por cada bloque. Pero los filtros son pequenos (unos 20 KB por bloque en promedio) y pueden verificarse contra un compromiso de hash en la cadena de cabeceras. El cliente no necesita confiar en que el servidor proporcione filtros precisos.

Ejecutar Bitcoin sobre **Tor** agrega otra capa de privacidad. Tor oculta la direccion IP del cliente a los nodos a los que se conecta. Combinado con Compact Block Filters, un cliente ligero puede verificar su saldo sin revelar su identidad ni sus direcciones a ningun par.
