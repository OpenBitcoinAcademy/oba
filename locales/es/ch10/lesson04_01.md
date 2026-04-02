## Clientes SPV

No todos los dispositivos pueden almacenar y validar la blockchain completa. Los telefonos moviles, dispositivos embebidos y billeteras de hardware tienen almacenamiento y poder de procesamiento limitados. Estos dispositivos ejecutan clientes **SPV** (Simplified Payment Verification).

Un cliente SPV descarga solo las cabeceras de bloque, no bloques completos. La cadena de cabeceras es pequena: unos 60 MB para toda la historia de Bitcoin. El cliente puede verificar que una cabecera de bloque cumple el objetivo de proof of work, confirmando que un minero gasto energia real para producirla.

Para comprobar si una transaccion esta confirmada, el cliente SPV solicita una **prueba de Merkle** a un nodo completo. La prueba muestra que el hash de la transaccion esta incluido en la raiz de Merkle de un bloque. El cliente confia en que los nodos completos validaron las transacciones del bloque, porque falsificar una cabecera de proof of work valida es prohibitivamente costoso.

SPV da a los clientes ligeros un nivel razonable de certeza sin descargar la blockchain completa. La contrapartida es que un cliente SPV no puede detectar si un bloque contiene una transaccion invalida. Confia en el peso economico del proof of work en lugar de verificar cada regla por si mismo.
