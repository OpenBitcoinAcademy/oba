## Privacidad mediante revelacion selectiva

Una salida Taproot con 8 hojas de script revela solo una hoja cuando se gasta por ruta de script. Las otras 7 hojas aparecen como hashes en la prueba de Merkle. Un observador conoce una condicion de gasto pero no puede determinar cuales eran las otras condiciones.

Compara esto con multisig P2WSH: gastar revela el script completo, todas las claves publicas y cuales claves firmaron. Todas las partes involucradas son visibles.

Un multisig 2-de-3 usando Taproot: la ruta de clave contiene el agregado MuSig2 de los dos firmantes mas comunes. Dos hojas de script contienen los pares de respaldo (A+C y B+C). En el caso comun, se usa la ruta de clave y nada sobre el multisig se revela. En un respaldo, solo un par alternativo queda expuesto. El otro par permanece como un hash.
