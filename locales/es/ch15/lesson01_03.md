## Lo que Miniscript resuelve

Miniscript, especificado en BIP 379, es una representacion estructurada de Bitcoin Script. Mapea un conjunto definido de fragmentos componibles a opcodes de Script. Cada fragmento tiene propiedades conocidas: su tipo, costo de recursos y datos de testigo requeridos.

Dado que el mapeo es estructurado, el software puede analizar una expresion Miniscript y determinar cada ruta de gasto, cada clave requerida, cada restriccion de timelock y el tamano exacto del testigo para cada ruta. Dos expresiones que codifican la misma politica pueden compararse y demostrarse equivalentes.

La composicion se vuelve segura. Un monedero puede tomar dos fragmentos Miniscript, combinarlos con AND u OR, y calcular el consumo de recursos del script resultante antes de transmitir. Si alguna ruta de ejecucion excede los limites de consenso, el compilador rechaza la composicion en tiempo de construccion, no en tiempo de gasto.
