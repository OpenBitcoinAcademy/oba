## Frases semilla BIP39

BIP39 codifica la semilla de una billetera como una secuencia de 12 o 24 palabras en inglés, llamada frase de recuperación (o frase semilla). Cada palabra proviene de una lista estandarizada de 2,048 palabras.

La codificación en palabras cumple dos propósitos. Las palabras son más fáciles de escribir, releer, y transcribir que el hexadecimal crudo. Y la última palabra incluye un checksum que detecta errores de transcripción.

Una frase de 12 palabras codifica 128 bits de entropía. Una frase de 24 palabras codifica 256 bits. Ambas son lo suficientemente fuertes para las necesidades de seguridad actuales. La frase se convierte en una semilla de 512 bits usando PBKDF2 con 2,048 rondas de HMAC-SHA512.
