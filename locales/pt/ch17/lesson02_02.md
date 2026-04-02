## Interpolacao de Lagrange

FROST se baseia no Compartilhamento de Segredo de Shamir. Durante a geracao de chaves, a chave secreta do grupo e dividida em n shares de forma que quaisquer t shares possam reconstrui-la. A ferramenta matematica que torna isso possivel e a interpolacao de Lagrange.

Um polinomio de grau t-1 e determinado de forma unica por t pontos. Cada assinante detem um ponto (seu share secreto) em um polinomio cujo termo constante e o segredo do grupo. Durante a assinatura, os assinantes nunca reconstroem o polinomio completo. Em vez disso, cada assinante multiplica seu share de assinatura por um coeficiente de Lagrange que depende do conjunto de assinantes participantes.

O coeficiente de Lagrange para o assinante $i$ em um conjunto de assinatura $S$ e: $\lambda_i = \prod_{j \in S, j \neq i} \frac{j}{j - i}$. Esses coeficientes garantem que a soma dos shares de assinatura ponderados produza uma assinatura Schnorr valida para a chave publica do grupo, sem que nenhuma parte segure ou reconstrua a chave secreta completa.
