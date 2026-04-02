## Sincronizar la blockchain

Un nodo nuevo debe descargar y validar toda la blockchain antes de poder verificar transacciones actuales. Este proceso se llama **Descarga Inicial de Bloques** (IBD, por sus siglas en ingles).

El nodo envia mensajes `getheaders` a sus pares, solicitando cabeceras de bloque en lotes. Las cabeceras son pequenas (80 bytes cada una) y llegan rapidamente. El nodo las usa para construir la cadena de cabeceras con el mayor proof of work acumulado.

Una vez que el nodo identifica la mejor cadena de cabeceras, solicita bloques completos usando mensajes `getdata`. Descarga bloques de multiples pares en paralelo para acelerar el proceso. Cada bloque se valida contra las reglas de consenso al llegar: firmas de transacciones, ejecucion de scripts, limites de peso y el objetivo de proof of work.

IBD puede tomar horas o dias en hardware lento. Despues de completarse, el nodo cambia a operacion normal. Recibe nuevos bloques y transacciones a medida que se transmiten y los valida en tiempo real.
