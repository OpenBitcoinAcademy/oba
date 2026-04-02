## Actualizador y Firmante

El Actualizador agrega los metadatos que los Firmantes necesitan. Para cada entrada, el Actualizador adjunta el UTXO que se gasta (ya sea la transaccion previa completa o el UTXO de testigo especifico), la ruta de derivacion BIP 32 para la clave de firma, el redeem script si la entrada es P2SH, y el witness script si la entrada es P2WSH. Para cada salida, el Actualizador puede adjuntar rutas de derivacion BIP 32 para que el firmante pueda verificar que las salidas de cambio pertenecen al mismo monedero.

El Firmante lee el PSBT, identifica las entradas que puede firmar y produce firmas parciales. Para cada entrada que firma, escribe una entrada PARTIAL_SIG con la clave publica como identificador. El Firmante no modifica la transaccion en si. No finaliza scripts. Agrega su firma y pasa el PSBT al siguiente participante.

Un Firmante debe verificar los datos del UTXO antes de firmar. Si el PSBT afirma que una entrada gasta 0.5 BTC pero el UTXO real contiene 1.0 BTC, el firmante donaria 0.5 BTC en comisiones sin saberlo. Los monederos de hardware comparan los montos de UTXO con los montos de salida y advierten al usuario si la comision parece excesiva.
