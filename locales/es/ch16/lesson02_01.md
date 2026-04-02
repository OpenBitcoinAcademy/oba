## Roles, no personas

El flujo de trabajo de PSBT define siete roles: Creador, Actualizador, Firmante, Combinador, Finalizador, Extractor y (en PSBTv2) Constructor. Cada rol realiza un paso. Una sola persona o programa puede cumplir multiples roles, pero la separacion importa porque cada rol tiene diferentes requisitos de confianza y diferente acceso a la informacion.

El Creador construye el PSBT inicial. En v0 (BIP 174), el Creador produce la transaccion sin firmar y la envuelve en un PSBT con mapas de entrada y salida vacios. En v2 (BIP 370), el Creador establece campos globales como la version de la transaccion y el locktime, pero aun no incluye entradas ni salidas. Esa tarea recae en el Constructor.

El rol de Constructor existe solo en PSBTv2. Agrega entradas y salidas al PSBT. Multiples Constructores pueden colaborar: uno agrega las entradas que controla, otro agrega las suyas, y cada uno agrega las salidas que necesita. Esto permite la construccion interactiva de transacciones donde ninguna parte conoce la forma completa de la transaccion hasta que todos los Constructores han contribuido.
