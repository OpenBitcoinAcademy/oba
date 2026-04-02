## Estructura del Encabezado de Bloque

Cada bloque comienza con un encabezado de 80 bytes que contiene seis campos.

**Version** (4 bytes): indica que reglas de consenso sigue el bloque y senala soporte para propuestas de soft fork mediante versionbits.

**Hash del Bloque Anterior** (32 bytes): el hash doble SHA-256 del encabezado del bloque anterior. Esto enlaza cada bloque con su predecesor, formando la cadena.

**Raiz Merkle** (32 bytes): un hash unico que compromete a cada transaccion en el bloque. Cambiar cualquier transaccion cambia la raiz Merkle.

**Marca de Tiempo** (4 bytes): el tiempo aproximado en que se mino el bloque, en segundos de epoca Unix.

**Target Bits** (4 bytes): una codificacion compacta del objetivo de prueba de trabajo. El hash del encabezado del bloque debe estar por debajo de este objetivo.

**Nonce** (4 bytes): el valor que los mineros cambian para buscar un hash valido.
