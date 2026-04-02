## El algoritmo de firma de SegWit

Segwit introdujo un nuevo algoritmo de firma (BIP143) que corrige un problema de larga data: en las transacciones legacy, la cantidad que se gasta no se incluye en los datos firmados. Esto obligaba a los firmantes a obtener la transacción anterior completa para verificar la cantidad, ralentizando las billeteras de hardware.

BIP143 incluye la cantidad de cada entrada en los datos que se firman. Una billetera de hardware puede verificar el valor total de las entradas y la comisión sin descargar transacciones anteriores. El proceso de firma es más rápido y más seguro.

Para segwit v1 (Taproot), BIP341 define un algoritmo actualizado que incluye compromisos adicionales, soportando tanto el gasto por key path como por script path.
