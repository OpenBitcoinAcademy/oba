## Claves publicas X-only

Taproot usa claves publicas X-only de 32 bytes en lugar del formato comprimido tradicional de 33 bytes. La coordenada Y se elige implicitamente como par. Esto ahorra un byte por clave en la cadena y en las firmas.

Una clave publica comprimida estandar tiene un byte de prefijo (0x02 para Y par, 0x03 para Y impar) seguido de 32 bytes de coordenada X. Taproot elimina el prefijo y requiere que la coordenada Y sea par. Si la Q calculada tiene Y impar, la implementacion niega la clave privada (lo que cambia Y a par) antes de firmar.

Esta convencion simplifica la verificacion por lotes y la agregacion de claves. Toda clave de salida Taproot mide 32 bytes. Toda firma Schnorr mide 64 bytes. No hay codificaciones de longitud variable.
