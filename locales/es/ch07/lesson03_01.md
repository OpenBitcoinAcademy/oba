## SegWit Nativo (P2WPKH y P2WSH)

Segregated Witness movio los datos de firma fuera del script de entrada hacia una estructura de testigo separada. Esto corrigio la maleabilidad de transacciones (terceros podian modificar el txid alterando la codificacion de la firma) y habilito el descuento de testigo para las comisiones.

P2WPKH (Pay to Witness Public Key Hash) es el equivalente segwit de P2PKH. El script de salida contiene: OP_0 <hash-de-clave-publica-de-20-bytes>. El testigo proporciona la firma y la clave publica. El script de entrada esta vacio.

P2WSH (Pay to Witness Script Hash) es el equivalente segwit de P2SH. El script de salida contiene: OP_0 <hash-de-script-de-32-bytes>. El testigo proporciona el script y los datos requeridos. P2WSH usa un hash SHA-256 de 32 bytes en lugar del HASH160 de 20 bytes de P2SH, proporcionando mayor resistencia a colisiones.
