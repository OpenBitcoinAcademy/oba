## Filtros Bloom y su problema de privacidad

**BIP 37** introdujo los filtros Bloom como una forma para que los clientes SPV soliciten solo las transacciones relevantes a los nodos completos. Un filtro Bloom es una estructura de datos probabilistica. El cliente crea un filtro que contiene sus direcciones y lo envia a un par. El par envia solo las transacciones que coinciden con el filtro.

El diseno tiene un defecto serio de privacidad. El nodo completo ve el filtro Bloom y puede deducir que direcciones pertenecen al cliente. Aunque un filtro Bloom tiene falsos positivos, la investigacion ha demostrado que un nodo malicioso puede identificar las direcciones del cliente con alta precision analizando el patron de bits del filtro.

Un nodo que sirve solicitudes de filtros Bloom tambien soporta un costo computacional significativo. Debe probar cada transaccion en cada bloque contra el filtro. Esto crea un vector de denegacion de servicio: un atacante puede conectarse muchas veces con diferentes filtros y forzar al nodo a realizar trabajo costoso. Muchos operadores de nodos ahora desactivan el soporte de BIP 37.
