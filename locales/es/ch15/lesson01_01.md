## Escribir Bitcoin Script a mano

Bitcoin Script es un lenguaje basado en pila. Cada condicion de gasto es una secuencia de opcodes: insertar una clave, verificar una firma, comprobar un hash, aplicar un timelock. Para un gasto con una sola clave, el script es corto y bien entendido.

Al combinar condiciones, la dificultad crece rapido. Un multisig 2-de-3 con un respaldo por timelock requiere un orden cuidadoso de ramas OP_IF, OP_CHECKMULTISIG y OP_CHECKSEQUENCEVERIFY. Un opcode mal colocado cambia las condiciones de gasto por completo. Una clave duplicada puede pasar desapercibida hasta que los fondos quedan bloqueados.

Ninguna herramienta puede analizar un Script arbitrario y responder: "Quien puede gastar esto, bajo que condiciones?" El lenguaje carece de estructura. Diferentes secuencias de opcodes pueden codificar la misma politica, y no existe un metodo general para compararlas.
