## Scripts de Firma Multiple

Un script de firma multiple (multisig) requiere M firmas de N claves publicas para autorizar un gasto. Un multisig 2-de-3 requiere cualquier 2 de 3 claves designadas.

En una salida multisig directa, el script de bloqueo contiene: M <pubkey1> <pubkey2> ... <pubkeyN> N OP_CHECKMULTISIG. El script de desbloqueo proporciona: OP_0 <sig1> <sig2>. El OP_0 inicial es un parche para un error historico de desfase por uno en OP_CHECKMULTISIG.

En la practica, multisig se envuelve en P2SH. El remitente ve una direccion P2SH estandar. Los detalles del multisig estan ocultos dentro del redeem script y se revelan solo al momento de gastar. Esto mantiene las salidas compactas y la complejidad privada hasta el momento del gasto.
