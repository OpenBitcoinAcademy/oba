## Nodos de archivo y nodos podados

Un nodo completo que conserva cada bloque desde el bloque genesis se llama **nodo de archivo**. Puede servir datos historicos de bloques a cualquier par que los solicite. Los nodos nuevos que se unen a la red dependen de los nodos de archivo para descargar la blockchain completa.

Ejecutar un nodo de archivo requiere un espacio de disco considerable. A 2024, la blockchain supera los 600 GB. No todos los operadores de nodos completos pueden permitirse ese almacenamiento.

Un **nodo podado** valida cada bloque pero descarta los datos de bloques antiguos despues de procesarlos. Conserva solo el conjunto UTXO y los bloques mas recientes. Un nodo podado aplica todas las reglas de consenso. No depende de la confianza. La contrapartida: no puede servir bloques historicos a otros pares.

Tanto los nodos de archivo como los nodos podados son nodos completos. Ambos validan todo. La diferencia es si almacenan bloques antiguos para que otros los descarguen.
