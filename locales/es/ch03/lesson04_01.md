## Por que importan multiples implementaciones

Un protocolo definido por una sola implementacion es fragil. Si Bitcoin Core tiene un error, cada nodo que lo ejecuta se ve afectado simultaneamente. Un error en el codigo de consenso podria dividir la red o permitir transacciones invalidas.

Multiples implementaciones independientes del mismo protocolo fortalecen el ecosistema. Cuando dos implementaciones no coinciden sobre si un bloque es valido, el desacuerdo revela un error en una de ellas. Esto sucedio en 2013 cuando una diferencia entre BerkeleyDB y LevelDB causo una division de cadena no intencional. El incidente demostro tanto el riesgo del monocultivo como el valor de la diversidad.

Un protocolo sano es aquel que multiples equipos pueden implementar a partir de la especificacion. Cuantas mas implementaciones coincidan en cada caso limite, mayor confianza tiene la red en que las reglas son correctas y estan bien definidas.
