## Retransmision de bloques compactos

Definido en **BIP 152**, la retransmision de bloques compactos reduce el ancho de banda necesario para propagar un nuevo bloque. La idea clave: la mayoria de las transacciones en un nuevo bloque ya estan en el mempool del nodo receptor.

En lugar de enviar el bloque completo, el nodo anunciante envia un mensaje `cmpctblock`. Este contiene la cabecera del bloque, una lista de IDs de transaccion cortos y la transaccion coinbase. Los IDs cortos de transaccion son hashes truncados de 6 bytes que permiten al receptor emparejar las transacciones que ya tiene.

El nodo receptor reconstruye el bloque a partir de su propio mempool usando los IDs cortos. Si faltan transacciones, las solicita con un mensaje `getblocktxn` y las recibe en una respuesta `blocktxn`.

BIP 152 define dos modos. En **modo de bajo ancho de banda**, un nodo primero anuncia el bloque con un mensaje `inv`. El par solicita el bloque compacto solo si esta interesado. En **modo de alto ancho de banda**, el nodo envia el bloque compacto inmediatamente sin esperar una solicitud. Los mineros y nodos bien conectados suelen usar el modo de alto ancho de banda para minimizar la latencia.

La retransmision de bloques compactos reduce el ancho de banda de propagacion en un 90% o mas para un bloque tipico. Una propagacion mas rapida da a los mineros pequenos una oportunidad mas justa, fortaleciendo la descentralizacion.
