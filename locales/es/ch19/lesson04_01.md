## Ark: UTXOs fuera de la cadena

Ark es un protocolo de Capa 2 que permite transacciones de Bitcoin baratas sin requerir canales. Un Ark Service Provider (ASP) coordina rondas donde los usuarios intercambian UTXOs Virtuales (VTXOs): transacciones de Bitcoin validas mantenidas fuera de la cadena.

Cada VTXO es una hoja en un arbol de transacciones enraizado en un solo UTXO compartido en cadena. Los usuarios pueden transmitir su VTXO para reclamar fondos en cadena en cualquier momento (salida unilateral). Durante la operacion normal, el ASP agrupa miles de transferencias en una sola liquidacion en cadena.

Ark no requiere abrir canales ni bloquear liquidez por adelantado. Los usuarios reciben VTXOs al ceder los antiguos en rondas periodicas. Los intercambios atomicos mediante salidas conectoras aseguran que no se pierdan fondos durante el intercambio.
