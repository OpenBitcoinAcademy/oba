## Campos Taproot y diferencias de version

BIP 371 agrego campos especificos de Taproot al formato PSBT. TAP_KEY_SIG (tipo de clave de entrada 0x13) almacena una firma Schnorr para un gasto por ruta de clave. TAP_SCRIPT_SIG (tipo de clave de entrada 0x14) almacena una firma Schnorr para una hoja de script especifica, identificada tanto por la clave publica X-only como por el hash de la hoja.

TAP_LEAF_SCRIPT (tipo de clave de entrada 0x15) proporciona el script, su version de hoja y el bloque de control necesario para el gasto por ruta de script. TAP_BIP32_DERIVATION (tipo de clave de entrada 0x16) extiende el campo estandar de derivacion BIP 32 con una lista de hashes de hojas para las cuales la clave puede firmar. TAP_INTERNAL_KEY (tipo de clave de entrada 0x17) registra la clave publica interna antes del tweaking.

En el lado de las salidas, TAP_INTERNAL_KEY (tipo de clave de salida 0x05) y TAP_BIP32_DERIVATION (tipo de clave de salida 0x07) permiten a los firmantes verificar que las salidas de cambio Taproot pertenecen al mismo monedero. El firmante puede re-derivar la clave con tweak desde la clave interna y confirmar que coincide con el scriptPubKey de la salida.

PSBTv2 (BIP 370) difiere de v0 en estructura, no en concepto. En v0, la transaccion sin firmar vive en el mapa global como un blob serializado unico. En v2, las entradas y salidas se describen por campos por-mapa: PREVIOUS_TXID, OUTPUT_INDEX, SEQUENCE para entradas; AMOUNT y SCRIPT para salidas. Esto permite a los Constructores agregar entradas y salidas de forma incremental sin re-serializar toda la transaccion cada vez.
