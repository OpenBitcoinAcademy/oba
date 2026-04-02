## Encontrar pares

Cuando un nodo de Bitcoin se inicia por primera vez, no conoce otros nodos. Necesita encontrar al menos un par para comenzar a participar en la red.

Bitcoin usa **DNS seeds** para el descubrimiento inicial. Los DNS seeds son servidores DNS operados por miembros de la comunidad Bitcoin. Devuelven las direcciones IP de nodos completos conocidos y estables. Bitcoin Core tiene varios DNS seeds codificados en su codigo fuente.

Si DNS no esta disponible (bloqueado o censurado), Bitcoin Core tambien incluye una lista de direcciones IP codificadas como ultimo recurso. Estas direcciones se actualizan con cada version del software.

Una vez que un nodo se conecta a su primer par, puede pedirle direcciones adicionales. El nodo construye un conjunto diverso de conexiones, tipicamente 8 salientes y hasta 125 entrantes.
