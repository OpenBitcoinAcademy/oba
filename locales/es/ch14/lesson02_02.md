## MuSig2 como clave interna

La clave interna P puede ser un agregado MuSig2 de multiples claves publicas. Si Alice y Bob agregan sus claves en P usando MuSig2, la clave ajustada Q se compromete con ambos. Cuando cooperan, producen una sola firma Schnorr sobre Q.

En la cadena, este multisig 2-de-2 se ve identico a una transaccion de un solo firmante. La salida ocupa 34 bytes. El testigo ocupa 64 bytes. Ningun observador puede determinar que hubo dos partes involucradas.

BitGo reporto que una entrada de ruta de clave MuSig2 cuesta 57.5 bytes virtuales, comparado con 104.5 vbytes para una entrada multisig P2WSH nativa de SegWit. El ahorro proviene de eliminar las claves publicas y firmas individuales que los scripts multisig requieren en la cadena.
