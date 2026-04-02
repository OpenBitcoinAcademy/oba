## Tres capas: SimplPedPop, EncPedPop, ChillDKG

La base es SimplPedPop. Cada participante genera su propio polinomio aleatorio de grado t-1 y envia una evaluacion secreta a cada otro participante, junto con un compromiso de los coeficientes de su polinomio. Cada participante suma las evaluaciones que recibio para calcular su porcion secreta final. La clave publica del grupo se deriva de la suma de los compromisos de todos los participantes con sus terminos constantes. Ninguna parte posee el secreto completo.

SimplPedPop asume un canal seguro entre cada par de participantes. EncPedPop agrega esto haciendo que cada participante genere un par de claves de cifrado efimero y cifre sus evaluaciones secretas con la clave publica del destinatario. Ahora el protocolo funciona sobre un canal de difusion publico, porque los espias no pueden descifrar las porciones secretas.

ChillDKG envuelve EncPedPop con una capa de gestion de sesion. Introduce una clave secreta de host que cada participante mantiene de forma persistente, un conjunto de datos de recuperacion comun que es identico para todos los participantes, y un protocolo para detectar y manejar el mal comportamiento. La clave secreta de host, combinada con los datos de recuperacion comunes, permite a un participante reconstruir su porcion si pierde su dispositivo de firma.
