## De semillas a árboles

Una semilla es un número aleatorio, típicamente de 128 a 256 bits. Es la raíz de toda la derivación de claves. Desde esta semilla se puede generar una jerarquía de claves: la semilla produce una clave maestra, la clave maestra produce claves hijas, y cada hija puede producir sus propias hijas.

Esta estructura de árbol es la base de las billeteras Hierarchical Deterministic (HD), definidas en BIP32. El árbol permite organizar claves por propósito: una rama para recibir pagos, otra para cambio, otra para una cuenta diferente.

La semilla nunca cambia. Cada clave en el árbol se puede regenerar a partir de ella. Un solo respaldo protege un número ilimitado de claves futuras.
