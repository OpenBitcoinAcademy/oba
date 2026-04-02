## Geracao Deterministica de Chaves

A geracao deterministica de chaves resolve o problema de backup. Uma unica semente aleatoria produz uma sequencia de chaves privadas por meio de hashing repetido. Cada chave na sequencia pode ser recriada a partir da semente.

Faca backup da semente uma vez e voce podera recuperar todas as chaves que a carteira gerou ou gerara. A semente e o segredo-mestre. Perca-a e voce perde acesso a todas as chaves derivadas.

A forma mais simples: comece com uma semente, faca hash para obter a chave 1, faca hash da chave 1 para obter a chave 2, e assim por diante. Isso produz uma lista plana de chaves. Mas carteiras modernas usam uma abordagem mais sofisticada: uma arvore.
