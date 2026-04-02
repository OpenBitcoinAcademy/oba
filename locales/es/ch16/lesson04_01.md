## El flujo de trabajo con monedero de solo lectura

Un flujo de trabajo de firma con monedero de hardware comienza con un monedero de solo lectura en un ordenador conectado a la red. El monedero de solo lectura contiene las claves publicas extendidas (xpubs) de todos los firmantes pero ninguna clave privada. Puede generar direcciones, rastrear saldos y construir transacciones, pero no puede firmar.

Cuando el usuario quiere gastar, el monedero de solo lectura crea un PSBT. Completa la transaccion sin firmar, adjunta datos WITNESS_UTXO para cada entrada y escribe entradas BIP32_DERIVATION para que el monedero de hardware sepa que claves usar. El PSBT se exporta a un archivo en una tarjeta SD o se codifica como un codigo QR animado.

El monedero de hardware recibe el PSBT, analiza las entradas y salidas, y muestra un resumen al usuario: que direcciones reciben fondos, cuanto va a cada una y cuanto se paga en comisiones. El usuario confirma en el dispositivo. El monedero de hardware deriva las claves privadas de su semilla usando las rutas BIP 32 del PSBT, firma cada entrada que puede, escribe entradas PARTIAL_SIG y exporta el PSBT actualizado de vuelta por tarjeta SD o codigo QR.

El monedero de solo lectura importa el PSBT firmado. Si todas las firmas requeridas estan presentes, finaliza y extrae la transaccion cruda, luego la transmite. Si se necesitan mas firmas (como en una configuracion multisig), pasa el PSBT al siguiente firmante.
