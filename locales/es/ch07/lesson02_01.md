## Pay to Script Hash (P2SH)

P2SH separa el script de la direccion. En lugar de codificar el script de bloqueo completo en la salida, P2SH codifica solo un hash del script. El gastador revela el script completo (llamado redeem script) al momento de gastar.

El script de salida contiene: OP_HASH160 <script_hash> OP_EQUAL. Para gastar, la entrada proporciona el redeem script y los datos que el redeem script requiera (firmas, claves publicas). La red aplica hash al redeem script proporcionado y verifica que coincida con el hash en la salida.

Las direcciones P2SH comienzan con "3" en mainnet. Permiten condiciones de gasto complejas (multisig, timelocks) sin requerir que el remitente conozca los detalles. El remitente paga a un hash corto. El receptor maneja la complejidad.
