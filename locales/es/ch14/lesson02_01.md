## Gastar con la clave ajustada

El gasto por ruta de clave es la forma mas sencilla y privada de gastar una salida Taproot. El testigo contiene un solo elemento: una firma Schnorr BIP 340 de 64 bytes para la clave ajustada Q.

El gastador calcula la clave privada ajustada (clave privada original + ajuste), firma la transaccion y proporciona la firma. No se revelan scripts. No se necesitan pruebas de Merkle. El testigo completo ocupa 64 bytes.

Un verificador comprueba la firma contra Q (la clave en el scriptPubKey de la salida). Si es valida, el gasto esta autorizado. El verificador no sabe ni le importa si Q se derivo de una sola clave, un agregado MuSig2 o una clave con un arbol de scripts incrustado.
