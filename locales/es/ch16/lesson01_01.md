## El problema antes de los PSBTs

Una transaccion de Bitcoin necesita informacion de multiples fuentes antes de poder transmitirse. El monedero que crea la transaccion debe saber que UTXOs gastar. El firmante debe poseer las claves privadas. En muchas configuraciones, estos son dispositivos diferentes o software diferente.

Antes de BIP 174, no existia una forma estandar de pasar una transaccion sin firmar entre estos participantes. Cada software de monedero inventaba su propio formato. Bitcoin Core serializaba las transacciones parciales de forma diferente a Electrum, que las serializaba de forma diferente a los monederos de hardware. Una transaccion creada en una herramienta no podia ser firmada por otra sin codigo de integracion personalizado.

Esto hacia que las configuraciones multifirma fueran penosas. Cada firmante necesitaba software compatible. Coordinar entre un portatil, un monedero de hardware y una maquina de almacenamiento en frio requeria pasos manuales y conversiones de formato que introducian errores.
