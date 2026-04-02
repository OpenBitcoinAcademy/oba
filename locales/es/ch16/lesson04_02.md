## Paso de multisig y transporte

En una configuracion multisig, el PSBT viaja a traves de cada firmante en secuencia o se distribuye a todos los firmantes en paralelo. Considera un multisig 2-de-3: el coordinador crea el PSBT y lo envia al firmante A y al firmante B simultaneamente. Cada firmante agrega su firma parcial y devuelve el PSBT. El coordinador combina ambos PSBTs, finaliza y transmite.

El mecanismo de transporte no esta especificado por BIP 174. Los PSBTs pueden moverse en tarjetas SD, como codigos QR (usando la codificacion UR de BCR-2020-005), por NFC, por correo electronico, a traves de un servicio para compartir archivos o por cualquier otro canal. El formato es autocontenido. Cada firmante obtiene todo lo que necesita del propio PSBT, sin necesidad de un canal lateral.

PSBT es texto plano, no cifrado. Cualquiera que intercepte un PSBT puede ver los montos de la transaccion, las direcciones involucradas y las firmas parciales recopiladas hasta el momento. Esto es un problema de privacidad, no de seguridad. Un atacante que lee un PSBT no puede robar fondos porque el PSBT no contiene claves privadas. Pero el atacante aprende lo que el usuario esta gastando y a donde van los fondos. Para transacciones sensibles, el PSBT debe transportarse por un canal cifrado o en medios fisicos que el usuario controle.
