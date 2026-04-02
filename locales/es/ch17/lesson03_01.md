## El problema del intermediario de confianza

La forma mas sencilla de configurar las porciones de clave FROST es con un intermediario de confianza. Una parte genera la clave secreta del grupo, evalua el polinomio de Shamir en n puntos y distribuye una porcion a cada participante. Esto funciona pero crea un punto unico de fallo: el intermediario conoce la clave secreta completa. Si el intermediario se ve comprometido, todos los fondos del grupo estan en riesgo.

Un intermediario comprometido tambien puede distribuir porciones inconsistentes, dando a algunos participantes datos invalidos que haran que la firma falle de forma impredecible. Los participantes no tienen forma de verificar que sus porciones son correctas sin un protocolo que garantice la consistencia.

ChillDKG, especificado en BIP 445, elimina al intermediario de confianza. Es un protocolo de generacion distribuida de claves: el grupo produce conjuntamente las porciones de clave sin que ninguna parte posea o vea la clave secreta completa. El protocolo se construye en tres capas, cada una agregando una garantia especifica.
