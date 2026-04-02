## El saludo inicial

Dos nodos establecen una conexion con un saludo version/verack sobre TCP puerto 8333 (el puerto predeterminado de la red Bitcoin).

El nodo que se conecta envia un mensaje `version` que contiene su version de protocolo, altura de bloque, hora actual y los servicios que soporta. El nodo receptor responde con su propio mensaje `version`. Luego cada nodo envia un `verack` (confirmacion de version) para confirmar.

Despues de que el saludo se completa, los nodos pueden intercambiar datos. Si las versiones de protocolo son incompatibles, la conexion se cierra.

Los nodos comparten direcciones de pares conocidos usando mensajes `addr` y `getaddr`. Cuando un nodo aprende nuevas direcciones, las almacena y puede conectarse a ellas mas tarde. Este protocolo de chisme permite que la red crezca y se repare sin ningun directorio central.
