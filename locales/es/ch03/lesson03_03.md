## El directorio de datos

Bitcoin Core almacena sus datos en un directorio especifico de la plataforma: `~/.bitcoin` en Linux, `~/Library/Application Support/Bitcoin` en macOS, `%APPDATA%\Bitcoin` en Windows.

El subdirectorio `blocks/` contiene los archivos de datos de bloques crudos. El directorio `chainstate/` almacena la base de datos UTXO, un almacen LevelDB de todas las salidas no gastadas. El archivo `mempool.dat` persiste el mempool entre reinicios.

El archivo de configuracion `bitcoin.conf` controla la configuracion de red, los limites de recursos, la autenticacion RPC y las opciones de funcionalidad. Opciones como `prune=550` habilitan el modo podado, y `txindex=1` construye un indice de todas las transacciones para consultas mas rapidas.
