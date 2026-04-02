## Enrutamiento basado en el origen

El remitente, no la red, elige la ruta de pago. Cada nodo Lightning mantiene una vista local de la topologia de la red (que nodos existen, que canales los conectan, sus capacidades y tarifas de comision). El remitente calcula una ruta a partir de esta informacion.

Esto difiere del enrutamiento de internet, donde cada router decide de forma independiente el siguiente salto. En Lightning, el remitente tiene control total sobre la ruta. Los nodos intermedios siguen instrucciones de reenvio sin conocer la ruta completa.
