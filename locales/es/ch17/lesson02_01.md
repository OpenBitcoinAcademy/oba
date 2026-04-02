## Firma en dos rondas

La firma FROST ocurre en dos rondas entre los firmantes participantes. Un coordinador retransmite los mensajes pero nunca aprende lo suficiente para falsificar una firma o extraer porciones secretas. El coordinador no es de confianza.

En la ronda uno, cada firmante genera un par de nonces aleatorios y envia los compromisos de nonce correspondientes al coordinador. El coordinador recopila los compromisos de todos los t firmantes participantes y los transmite al grupo. Estos compromisos vinculan a cada firmante con sus nonces antes de que nadie vea los valores de otros firmantes, previniendo que un firmante deshonesto elija nonces de forma adaptiva para manipular el resultado.

En la ronda dos, cada firmante calcula el nonce agregado a partir de los compromisos recopilados, construye el hash de desafio Schnorr y produce una porcion de firma usando su porcion de clave secreta y su nonce. Cada firmante envia su porcion de firma al coordinador.
