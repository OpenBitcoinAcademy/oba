## El sistema de tipos

Cada fragmento Miniscript tiene un tipo que describe su comportamiento en la pila de Script. Los cuatro tipos base son B, V, K y W.

Una expresion de tipo B (base) coloca un valor no-cero en caso de exito y cero en caso de fallo. Una expresion de tipo V (verify) tiene exito sin dejar nada en la pila o aborta el script. Una expresion de tipo K (key) coloca una clave publica en caso de exito. Una expresion de tipo W (wrapped) se comporta como B pero consume primero el elemento superior de la pila, usada para combinar sub-expresiones.

El compilador verifica los tipos en cada punto de composicion. Un fragmento `and_v` requiere que su primer argumento sea de tipo V y su segundo de tipo B. Si pasas argumentos con tipos incorrectos, el compilador rechaza la expresion. Esto captura errores que, en Script crudo, producirian un script que "funciona" pero aplica las condiciones equivocadas.

El sistema de tipos tambien rastrea propiedades como la disipabilidad (puede una rama insatisfecha dejar la pila limpia?) y la no-maleabilidad (existe exactamente un testigo valido para cada ruta de gasto?). Estas propiedades se propagan a traves de la composicion. Una expresion Miniscript las tiene o no, y el compilador puede informar cuales.
