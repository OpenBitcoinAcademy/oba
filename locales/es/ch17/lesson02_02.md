## Interpolacion de Lagrange

FROST se construye sobre el Secreto Compartido de Shamir. Durante la generacion de claves, la clave secreta del grupo se divide en n porciones tales que cualquier t porciones pueden reconstruirla. La herramienta matematica que hace esto posible es la interpolacion de Lagrange.

Un polinomio de grado t-1 queda determinado de forma unica por t puntos. Cada firmante posee un punto (su porcion secreta) en un polinomio cuyo termino constante es el secreto del grupo. Durante la firma, los firmantes nunca reconstruyen el polinomio completo. En su lugar, cada firmante multiplica su porcion de firma por un coeficiente de Lagrange que depende del conjunto de firmantes participantes.

El coeficiente de Lagrange para el firmante $i$ en un conjunto de firma $S$ es: $\lambda_i = \prod_{j \in S, j \neq i} \frac{j}{j - i}$. Estos coeficientes aseguran que la suma de las porciones de firma ponderadas produce una firma Schnorr valida para la clave publica del grupo, sin que ninguna parte posea o reconstruya la clave secreta completa.
