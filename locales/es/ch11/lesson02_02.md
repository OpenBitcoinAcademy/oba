## La Cadena de Hashes

Cada encabezado de bloque contiene el hash del encabezado del bloque anterior. El bloque 500,000 contiene el hash del bloque 499,999, que contiene el hash del bloque 499,998, y asi sucesivamente hasta el bloque genesis (bloque 0).

Cambiar cualquier dato en el bloque 499,999 cambiaria su hash. El bloque 500,000 entonces contendria un hash de bloque anterior incorrecto y se volveria invalido. Para alterar un bloque historico, un atacante debe rehacer la prueba de trabajo de ese bloque y de todos los bloques posteriores.

Este enlace acumulativo es lo que hace a la blockchain resistente a alteraciones. Mientras mas profundo esta un bloque, mas trabajo se requiere para alterarlo. Cada nuevo bloque agrega otra capa de proteccion a todos los bloques debajo de el.
