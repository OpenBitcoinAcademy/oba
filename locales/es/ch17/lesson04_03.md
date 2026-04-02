## El panorama general

Las firmas umbral representan un cambio en como funciona la custodia de Bitcoin. El enfoque legacy revelaba las politicas de gasto en cadena: cualquiera podia contar las claves, ver el umbral y rastrear monederos multisig a traves de transacciones. FROST y ChillDKG mueven toda esa complejidad fuera de la cadena.

El grupo negocia su umbral y genera claves usando ChillDKG. Cualquier t firmantes produce una firma usando el protocolo FROST. La blockchain ve un gasto por ruta de clave Taproot. Los verificadores comprueban una firma contra una clave. El costo en cadena es constante sin importar cuantos participantes esten involucrados: 64 bytes para la firma, 32 bytes para la clave.

Esto es posible porque Taproot y la verificacion Schnorr de BIP 340 ya estan desplegados en la red Bitcoin. No se requiere un soft fork. No se necesitan nuevos opcodes. El protocolo de firma umbral se ejecuta completamente en los monederos de los participantes. La capa de consenso verifica el resultado usando las mismas reglas que usa para cualquier otro gasto Taproot.
