## Ampliando el multisig con el tiempo

Las politicas con timelock pueden ampliar su conjunto de firmantes con el tiempo. Un negocio podria comenzar con una politica 2-de-2 entre dos cofundadores. Despues de un ano, una tercera clave (en manos de un asesor legal) se activa, y la politica se convierte en 2-de-3.

La expresion Miniscript es: `or(multi(2,founder_a,founder_b),and(multi(2,founder_a,founder_b,counsel),older(52560)))`. Durante el primer ano, solo los dos fundadores pueden firmar. Despues de 52,560 bloques (aproximadamente un ano), dos de las tres claves cualesquiera pueden firmar.

La politica completa vive en un solo UTXO. No se necesita ninguna transaccion en cadena cuando expira el timelock. La ruta de gasto adicional fue comprometida en el momento de creacion de la salida. La clave del asesor gana poder de gasto automaticamente cuando la altura de bloque supera el umbral. Esto elimina la necesidad de una ceremonia activa de claves en el punto de transicion.

Miniscript hace estas composiciones auditables. El compilador puede enumerar cada ruta de gasto y sus condiciones. Un revisor puede verificar que la clave del asesor no tiene poder de gasto antes del timelock, que los fundadores mantienen control total en todo momento, y que los tamanos de testigo se mantienen dentro de los limites de consenso.
