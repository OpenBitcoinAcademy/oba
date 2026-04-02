## Campos por entrada

Cada mapa de entrada describe una entrada de la transaccion. Los campos criticos que un firmante necesita son los datos del UTXO, la ruta de derivacion de la clave de firma y cualquier script requerido para construir el testigo.

WITNESS_UTXO (tipo de clave 0x01) almacena la salida que se gasta: el monto en satoshis y el scriptPubKey. Para entradas SegWit, esto es suficiente para firmar porque el algoritmo de sighash compromete el monto directamente. NON_WITNESS_UTXO (tipo de clave 0x00) almacena la transaccion previa completa, necesaria para entradas legacy donde el firmante debe verificar el monto buscando la salida en la transaccion completa.

BIP32_DERIVATION (tipo de clave 0x06) mapea una clave publica a su ruta de derivacion BIP 32 y la huella de la clave maestra. El firmante compara la huella con su propia clave maestra, deriva la clave privada en la ruta dada y firma. PARTIAL_SIG (tipo de clave 0x02) almacena una firma identificada por la clave publica que la produjo. SIGHASH_TYPE (tipo de clave 0x03) especifica que flag de sighash debe usar el firmante.

Para entradas P2SH, REDEEM_SCRIPT (tipo de clave 0x04) lleva el redeem script. Para entradas P2WSH, WITNESS_SCRIPT (tipo de clave 0x05) lleva el witness script. El firmante necesita estos para calcular el sighash correcto.
