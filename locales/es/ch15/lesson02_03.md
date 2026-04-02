## Mapeo biyectivo

La relacion entre Miniscript y Script es biyectiva para el conjunto de fragmentos soportados. Cada expresion Miniscript se mapea a exactamente una codificacion en Script, y cada patron de Script soportado se mapea de vuelta a exactamente una expresion Miniscript.

Considera el fragmento `pk(K)`: se codifica como `<K> OP_CHECKSIG`. El fragmento `and_v(v:pk(A),pk(B))` se codifica como `<A> OP_CHECKSIGVERIFY <B> OP_CHECKSIG`. No hay ambiguedad. Dado el Script, puedes recuperar el Miniscript original.

Esta propiedad hace posible el analisis de ida y vuelta. Un monedero puede recibir un Script de otra parte, decodificarlo a Miniscript, y determinar las condiciones de gasto sin confiar en la descripcion del remitente. Dos monederos construidos por equipos diferentes pueden acordar una politica, compilar a Miniscript de forma independiente, y verificar que produjeron el mismo Script.

No todos los Bitcoin Scripts validos pueden representarse en Miniscript. El conjunto de fragmentos cubre los patrones necesarios para politicas de gasto practicas: claves, hashes, timelocks, umbrales y combinaciones booleanas. Los Scripts que usan secuencias inusuales de opcodes o manipulacion de pila quedan fuera del subconjunto de Miniscript.
