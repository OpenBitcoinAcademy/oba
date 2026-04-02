## Tipos de nodos

No todos los nodos ejecutan el mismo software ni cumplen la misma funcion. Los nodos difieren en los datos que almacenan y las funciones que proporcionan.

Un **nodo completo** descarga y valida cada bloque y cada transaccion contra las reglas de consenso. No confia en nadie. Puede verificar cualquier pago de forma independiente. Bitcoin Core es la implementacion de nodo completo mas comun.

Un **nodo minero** es un nodo completo que tambien compite para crear nuevos bloques. Ensambla bloques candidatos a partir del mempool y realiza proof of work. Los nodos mineros ganan el subsidio de bloque y las comisiones de transaccion cuando encuentran un bloque valido.

Un **nodo ligero** (tambien llamado cliente SPV) no descarga bloques completos. Descarga solo las cabeceras de bloque y solicita pruebas de que transacciones especificas existen. Los nodos ligeros confian en los nodos completos hasta cierto punto, sacrificando seguridad a cambio de menor uso de recursos.
