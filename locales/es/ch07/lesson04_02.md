## MAST y Tapscript

El arbol Merkle en Taproot es un Merklized Alternative Script Tree (MAST). Cada hoja del arbol es un script diferente (una condicion de gasto diferente). Para gastar por ruta de script, el gastador revela solo el script que esta usando y proporciona una prueba Merkle de que es parte del arbol comprometido.

Los scripts no utilizados permanecen ocultos. Una salida Taproot con 100 posibles condiciones de gasto se ve igual en la blockchain que una con 1 condicion, a menos que se ejerza la ruta de script. Las ramas no utilizadas permanecen privadas.

Tapscript (BIP342) es el lenguaje de scripts usado dentro de las hojas de ruta de script de Taproot. Es similar al Script tradicional pero con mejoras: OP_CHECKSIGADD reemplaza a OP_CHECKMULTISIG (habilitando validacion por lotes), los opcodes de firma usan Schnorr en lugar de ECDSA, y el limite de tamano de script se elimina (reemplazado por limites de peso).
