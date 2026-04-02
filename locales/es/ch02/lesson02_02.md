## Scripts de desbloqueo

Para gastar una salida bloqueada, el gastador proporciona un script de desbloqueo, llamado scriptSig. Para P2PKH, el script de desbloqueo contiene una firma digital y la clave pública del gastador.

Bitcoin ejecuta estos scripts en una máquina de pila. Primero, el script de desbloqueo se ejecuta y coloca datos en la pila. Luego la pila se copia, y el script de bloqueo se ejecuta contra esa pila copiada. Los dos scripts nunca se combinan en uno solo. Esta separación se introdujo en 2010 para corregir una vulnerabilidad de seguridad.

Si el script de bloqueo termina con un valor verdadero en la parte superior de la pila, el gasto es válido. Si la pila está vacía o el valor superior es cero, el gasto falla.
