## Como se propagan los bloques

Cuando un minero encuentra un bloque valido, lo transmite a sus pares. Esos pares lo validan y lo reenvian a sus propios pares. El bloque se extiende por la red en pocos segundos.

La velocidad importa. Si dos mineros encuentran bloques validos casi al mismo tiempo, el bloque que alcanza a mas nodos primero tiene ventaja. Los mineros que se enteran del bloque A primero construyen sobre el. Los mineros que se enteran del bloque B primero construyen sobre ese. La carrera se resuelve cuando el siguiente bloque extiende una cadena, y el otro bloque queda obsoleto.

La propagacion lenta perjudica la descentralizacion. Si los grandes pools de mineria con conexiones rapidas se enteran de los bloques primero, tienen ventaja para el siguiente bloque. Los mineros mas pequenos con conexiones lentas desperdician trabajo en bloques obsoletos con mas frecuencia. El protocolo de red aborda esta asimetria mediante la retransmision de bloques compactos.
