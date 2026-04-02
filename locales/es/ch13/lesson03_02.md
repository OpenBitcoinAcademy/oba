## Riesgos de consenso

Las aplicaciones que validan transacciones deben implementar las reglas de consenso de forma exacta. Una diferencia sutil en como dos implementaciones manejan un caso limite puede causar que no esten de acuerdo sobre la validez de un bloque, dividiendo la red desde su perspectiva.

El enfoque mas seguro para las aplicaciones: delegar la validacion de consenso a un nodo completo (Bitcoin Core o equivalente) y usar su API para consultas de blockchain. No reimplementes las reglas de consenso en el codigo de la aplicacion a menos que estes preparado para coincidir con cada caso limite de la implementacion de referencia.

Para desarrolladores de billeteras: usa bibliotecas bien probadas para la derivacion de claves, generacion de direcciones y firma de transacciones. Prefiere implementaciones de codigo abierto probadas en combate sobre codigo personalizado.
