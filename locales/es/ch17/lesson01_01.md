## El panorama del multisig

Bitcoin ha soportado custodia multi-parte desde 2012 a traves de OP_CHECKMULTISIG. Un script multisig 2-de-3 lista las tres claves publicas en cadena y requiere dos firmas validas. Esto funciona, pero revela exactamente cuantas partes controlan los fondos. Cualquiera que inspeccione la blockchain ve el umbral, el numero total de firmantes y cada clave publica involucrada.

El costo en cadena escala linealmente. Un multisig 3-de-5 coloca cinco claves publicas de 33 bytes y tres firmas de 72 bytes en el testigo. Son 381 bytes de datos de testigo sin contar el script en si. Una configuracion 7-de-10 cuesta aun mas. Las comisiones suben, la privacidad baja y la politica de gasto es visible para todos los observadores.

Taproot y las firmas Schnorr cambiaron lo que es posible. MuSig2 permite que n participantes produzcan una sola firma agregada que luce identica a una firma Schnorr individual en cadena. La clave publica combinada ocupa 32 bytes. La firma ocupa 64 bytes. Ningun observador puede saber cuantas partes estuvieron involucradas.
