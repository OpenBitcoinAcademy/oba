## Creación y verificación

Crear una firma requiere dos entradas: la clave privada y el mensaje (los datos de la transacción, o un hash de ellos). El algoritmo de firma produce un valor de firma único tanto para la clave como para el mensaje.

Verificar una firma requiere tres entradas: la clave pública, el mensaje y la firma. El algoritmo de verificación produce verdadero o falso. Si es verdadero, la firma fue producida por el titular de la clave privada correspondiente para ese mensaje exacto.

La clave privada nunca se revela durante la creación ni la verificación. Cualquiera puede verificar, pero solo el titular de la clave puede firmar.
