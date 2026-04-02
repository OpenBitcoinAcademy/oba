## De n-de-n a t-de-n

MuSig2 tiene una limitacion: requiere que todos los n participantes firmen. Cada titular de clave debe estar en linea y cooperando al momento de firmar. Si un participante pierde su clave o se desconecta, los fondos quedan atrapados. Esto hace de MuSig2 un esquema puramente n-de-n.

FROST (Flexible Round-Optimized Schnorr Threshold signatures) resuelve esto. FROST es un esquema de firmas umbral: cualquier t de n participantes puede producir una firma valida. Una configuracion FROST 3-de-5 significa que tres cualesquiera de los cinco titulares de clave pueden firmar. Los otros dos pueden estar fuera de linea, no disponibles o incluso permanentemente perdidos.

La firma que FROST produce es una firma Schnorr estandar de 64 bytes. En cadena, es indistinguible de un gasto con una sola clave. La politica umbral, el numero de firmantes y las claves publicas individuales estan todos ocultos. Un gasto FROST luce exactamente como un gasto por ruta de clave desde una direccion Taproot individual.
