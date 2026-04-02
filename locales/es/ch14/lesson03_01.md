## Cuando la ruta de clave falla

Si la ruta de clave no puede usarse (un firmante no esta disponible, un timelock necesita ejercerse), la ruta de script proporciona una alternativa. El gastador revela un script del arbol comprometido en Q, demuestra que fue comprometido al momento de la creacion y satisface ese script.

El testigo para un gasto por ruta de script contiene: los datos que satisfacen el script (firmas, preimagenes), el script en si y un bloque de control. El bloque de control contiene la clave interna P, un byte de version de hoja y la prueba de Merkle (los hashes hermanos a lo largo del camino desde la hoja del script hasta la raiz).

El verificador reconstruye la raiz de Merkle a partir del script revelado y la prueba, calcula el ajuste esperado y verifica que Q sea igual a P + ajuste * G. Si la matematica es correcta, el script fue comprometido al momento de la creacion.
