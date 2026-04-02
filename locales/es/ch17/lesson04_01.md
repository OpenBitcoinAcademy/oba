## La reutilizacion de nonces es catastrofica

FROST hereda la sensibilidad de Schnorr a la gestion de nonces. Si un firmante usa el mismo nonce para dos sesiones de firma diferentes, un atacante que observe ambas porciones de firma puede extraer la porcion de clave secreta de ese firmante. Con suficientes porciones extraidas (t de ellas), el atacante puede reconstruir la clave secreta completa del grupo y robar todos los fondos.

Esto no es una preocupacion teorica. La generacion deterministica de nonces, que es segura para Schnorr con un solo firmante, es peligrosa en el entorno umbral. Si un coordinador malicioso repite una solicitud de firma antigua, un firmante que use nonces deterministicos produciria una nueva porcion de firma con el mismo nonce, filtrando su porcion. FROST por lo tanto requiere nonces criptograficamente aleatorios para cada sesion de firma.

Las implementaciones deben generar nonces de una fuente aleatoria fuerte y nunca deben persistir nonces entre sesiones. Si una sesion de firma se aborta, los nonces usados en esa sesion deben descartarse. Re-entrar a una sesion con los mismos nonces equivale a reutilizar nonces.
