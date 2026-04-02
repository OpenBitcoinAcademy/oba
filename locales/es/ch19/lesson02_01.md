## Sellos de un solo uso

RGB define un "sello" como un UTXO de Bitcoin. Cuando ocurre una transicion de estado (transferir un token, actualizar un contrato), el sello se "cierra" gastando ese UTXO. Un UTXO solo puede gastarse una vez, por lo que un sello solo puede cerrarse una vez. Esto previene el doble gasto de activos RGB usando el mecanismo de consenso existente de Bitcoin.

Los datos de la transicion de estado (quien posee que, actualizaciones del contrato) nunca tocan la blockchain. Solo un compromiso criptografico con la transicion se incrusta en la transaccion de Bitcoin, tipicamente en una salida Taproot. La blockchain ancla la temporalidad y el orden. Los datos viven fuera de la cadena con las partes involucradas.
