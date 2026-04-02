## Abortos identificables y MuSig2 vs FROST

Un participante malicioso puede interrumpir la firma enviando una porcion de firma invalida. Sin verificaciones adicionales, el coordinador sumaria las porciones y produciria una firma final invalida, pero no sabria cual firmante se comporto mal. FROST soporta abortos identificables: el coordinador puede verificar cada porcion de firma individualmente contra la porcion de clave publica del firmante. El firmante que se comporto mal es identificado y puede ser excluido de futuras sesiones.

MuSig2 y FROST sirven necesidades diferentes. MuSig2 es un esquema n-de-n: todos los participantes deben firmar, no hay umbral, y el protocolo es mas sencillo. FROST es un esquema t-de-n: tolera firmantes ausentes pero requiere una fase de generacion de claves mas compleja. Ambos producen una salida en cadena identica: una unica firma Schnorr de 64 bytes bajo una clave publica de 32 bytes.

MuSig2 es adecuado para escenarios donde se espera que todos los firmantes esten disponibles, como un canal Lightning entre dos nodos. FROST es adecuado para arreglos de custodia donde la redundancia importa, como una tesoreria corporativa con cinco titulares de clave donde tres cualesquiera pueden autorizar un gasto.
