## BIP32: El estándar de billeteras HD

BIP32 define cómo derivar un árbol de claves a partir de una sola semilla. El proceso usa HMAC-SHA512 para dividir cada paso de derivación en una clave hija y un chain code. El chain code se mezcla en la siguiente derivación, impidiendo que alguien calcule claves hermanas.

La clave maestra se ubica en la raíz del árbol. A partir de ella se deriva una secuencia de claves hijas, cada una identificada por un número de índice (0, 1, 2, ...). Cada hija puede producir sus propias hijas, formando un árbol de profundidad arbitraria.

Esta estructura de árbol le da a las billeteras poder organizativo. Distintas ramas del árbol sirven para distintos propósitos, y el árbol completo es recuperable a partir de la semilla original.
