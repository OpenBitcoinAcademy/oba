## Descriptores y Miniscript juntos

Los descriptores soportan la integracion de expresiones Miniscript. El descriptor `wsh()` envuelve una expresion Miniscript en una salida P2WSH. El descriptor `tr()` coloca una expresion Miniscript dentro de una hoja del arbol de scripts Taproot.

Un descriptor como `wsh(and_v(v:pk(Alice),or_d(pk(Bob),older(12960))))` define una politica de gasto completa en una sola cadena. Cualquier monedero que entienda el formato de descriptores puede importar esta cadena, derivar las direcciones correctas, identificar las condiciones de gasto y construir transacciones validas. El descriptor lleva la politica completa, no una pista parcial.

Esto es lo que hace a los monederos interoperables. Un firmante de hardware puede mostrar las condiciones de gasto analizadas del descriptor. Un monedero de solo lectura puede monitorear las direcciones resultantes. Un coordinador de firma puede construir un PSBT que satisfaga la politica. Cada herramienta lee la misma cadena de descriptor y llega a la misma salida de Script.

Los descriptores incluyen un checksum: ocho caracteres anexados despues de un simbolo `#`. El checksum detecta errores de transcripcion. Un monedero rechaza cualquier descriptor cuyo checksum no coincida. Un solo caracter cambiado en la clave o politica produce un checksum diferente, capturando errores antes de enviar fondos a la direccion equivocada.
