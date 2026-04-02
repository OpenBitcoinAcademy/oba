## Por que los monederos de hardware necesitaban esto

Los monederos de hardware tienen una restriccion especifica: no pueden consultar la blockchain. Reciben datos, verifican lo que pueden, firman y devuelven el resultado. Sin un formato estandar, cada fabricante de monederos de hardware tenia que escribir integraciones personalizadas para cada monedero de software. Agregar un nuevo monedero de hardware significaba parchear cada coordinador. Agregar un nuevo coordinador significaba parchear cada monedero de hardware. La explosion combinatoria era insostenible.

PSBT resolvio esto definiendo un formato que lleva toda la informacion que un firmante necesita. El UTXO que se gasta, la ruta de derivacion para la clave, el tipo de sighash a usar, el redeem script para P2SH, el witness script para P2WSH. Un monedero de hardware recibe un PSBT, lee los campos que necesita, firma, escribe su firma parcial de vuelta en el PSBT y lo devuelve. Sin protocolo propietario. Sin negociacion de formato.

El ecosistema convergio rapidamente. Coldcard, Trezor, Ledger, BitBox y Jade hablan PSBT. Los coordinadores de software como Sparrow, Specter y Bitcoin Core producen y consumen PSBTs. Un quorum multisig puede usar hardware de diferentes proveedores sin preocupaciones de compatibilidad.
