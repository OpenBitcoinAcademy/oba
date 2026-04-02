## Privacidad a traves de la indistinguibilidad

La ganancia de privacidad de las firmas umbral va mas alla de ocultar la politica de gasto. Con OP_CHECKMULTISIG, la plantilla del script en si es una huella digital. Las firmas de analisis de cadena categorizan direcciones por tipo de script, identificando monederos multisig e infiriendo arreglos de custodia.

FROST elimina esta huella digital. Un monedero de custodia 2-de-3, una tesoreria corporativa 5-de-7 y un monedero personal de una sola clave producen salidas en cadena identicas. Cada uno gasta con una clave publica de 32 bytes y una firma de 64 bytes dentro de una ruta de clave Taproot.

Esta indistinguibilidad beneficia a todos los usuarios de Taproot. Cuanto mayor es el conjunto de transacciones que lucen igual, mas dificil es distinguir una transaccion individual. Los usuarios de firmas umbral se mezclan con la multitud de gastadores de una sola clave, y los gastadores de una sola clave obtienen negacion plausible de que podrian ser firmantes umbral.
