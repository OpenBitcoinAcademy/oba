## Reajuste de dificultad

Bitcoin ajusta su dificultad cada 2,016 bloques, aproximadamente cada dos semanas. El objetivo es mantener un intervalo promedio de bloque de 10 minutos.

En cada punto de reajuste, los nodos calculan cuanto tiempo tomaron los 2,016 bloques anteriores. Si los bloques llegaron mas rapido que 10 minutos en promedio, el objetivo disminuye (la dificultad sube). Si los bloques llegaron mas lento, el objetivo aumenta (la dificultad baja). La formula compara el tiempo real transcurrido con los 20,160 minutos esperados.

Una proteccion impide que la dificultad cambie por mas de un factor de cuatro en un solo reajuste. Esto limita la rapidez con que la dificultad puede subir o caer. El mecanismo es completamente algoritmico. Ningun comite vota al respecto. Cada nodo calcula el mismo nuevo objetivo a partir de los mismos datos de la cadena.
