## Combinador, Finalizador, Extractor

El Combinador fusiona multiples PSBTs que contienen firmas parciales para la misma transaccion. En un multisig 2-de-3, el firmante A produce un PSBT con su firma y el firmante B produce otro. El Combinador toma ambos, fusiona las entradas PARTIAL_SIG de cada entrada y produce un solo PSBT con todas las firmas disponibles.

El Finalizador transforma las firmas parciales en un scriptSig y testigo completos para cada entrada. Para una entrada P2PKH, toma la unica PARTIAL_SIG y construye el scriptSig. Para un multisig P2WSH 2-de-3, toma las firmas parciales, las ordena correctamente y construye la pila de testigos con el redeem script. Tras la finalizacion, los mapas de entrada del PSBT contienen campos FINAL_SCRIPTSIG y FINAL_SCRIPTWITNESS. Los datos parciales ya no son necesarios.

El Extractor lee el PSBT finalizado y produce la transaccion de red cruda. Toma la transaccion sin firmar del mapa global, inserta los FINAL_SCRIPTSIG y FINAL_SCRIPTWITNESS de cada entrada y serializa el resultado. La salida es una transaccion de Bitcoin estandar lista para transmitir. El PSBT ha cumplido su proposito.
