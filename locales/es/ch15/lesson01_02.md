## Limites de recursos y fallos ocultos

El consenso de Bitcoin impone limites de recursos a los scripts: un tamano maximo de 10,000 bytes en legacy, un limite de 201 opcodes que no son push, y un presupuesto de sigops. Un script que exceda cualquier limite es invalido. La transaccion que lo contiene nunca se confirmara.

Al componer scripts a mano, el consumo de recursos es dificil de predecir. Cada rama de un arbol OP_IF contribuye de forma diferente al conteo de opcodes. Las condiciones anidadas multiplican la complejidad. Un script puede funcionar en una rama y exceder el limite de opcodes en otra, dependiendo de la ruta de ejecucion que tome el gastador.

Estos fallos son silenciosos. El script se ve correcto en un editor de texto. Puede pasar pruebas unitarias para una ruta de gasto. La ruta que falla se descubre solo cuando alguien intenta usarla en la red, y la transaccion es rechazada.
