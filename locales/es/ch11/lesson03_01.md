## Que es un Arbol Merkle?

Un arbol Merkle es un arbol binario de hashes. Cada nodo hoja es el hash de una transaccion. Cada nodo interno es el hash de la concatenacion de sus dos hijos. La raiz del arbol (la raiz Merkle) es un hash unico que compromete a cada transaccion en el bloque.

Si un bloque contiene cuatro transacciones (A, B, C, D), el arbol es: hash(A) y hash(B) se combinan en hash(AB). hash(C) y hash(D) se combinan en hash(CD). hash(AB) y hash(CD) se combinan en la raiz Merkle.

Si el numero de transacciones es impar, el ultimo hash se duplica para obtener un conteo par en cada nivel.
