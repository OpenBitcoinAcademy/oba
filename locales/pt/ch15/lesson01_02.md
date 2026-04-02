## Limites de Recursos e Falhas Ocultas

O consenso do Bitcoin aplica limites de recursos em scripts: tamanho maximo de 10.000 bytes no legacy, limite de 201 opcodes nao-push e orcamento de sigops. Um script que excede qualquer limite e invalido. A transacao que o carrega nunca confirmara.

Ao compor scripts a mao, o consumo de recursos e dificil de prever. Cada ramo de uma arvore OP_IF contribui de forma diferente para a contagem de opcodes. Condicoes aninhadas multiplicam a complexidade. Um script pode funcionar em um ramo e exceder o limite de opcodes em outro, dependendo de qual caminho de execucao o gastador toma.

Essas falhas sao silenciosas. O script parece correto em um editor de texto. Pode passar testes unitarios para um caminho de gasto. O caminho falho e descoberto apenas quando alguem tenta usa-lo na rede e a transacao e rejeitada.
