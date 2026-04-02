## Que es un PSBT

Una Transaccion de Bitcoin Parcialmente Firmada (PSBT) es un formato binario definido por BIP 174. Envuelve una transaccion sin firmar junto con los metadatos que cada participante necesita para cumplir su funcion. Los creadores adjuntan datos de UTXO. Los actualizadores agregan rutas de derivacion. Los firmantes agregan firmas. El formato lleva todo lo necesario para que cada rol pueda actuar de forma independiente.

El binario comienza con un magic de cinco bytes: `0x70736274FF`. Los primeros cuatro bytes deletrean "psbt" en ASCII. El quinto byte, `0xFF`, marca el separador. Cualquier herramienta que reciba un blob puede verificar estos cinco bytes para confirmar que es un PSBT antes de continuar con el analisis.

Despues del magic viene una secuencia de mapas clave-valor. Cada entrada del mapa tiene un tipo de clave, un payload de clave y un payload de valor. Un byte cero (0x00) termina cada mapa. El primer mapa es el mapa global. Los mapas posteriores alternan entre mapas por-entrada y por-salida, uno para cada entrada y salida de la transaccion.
