## O que o Miniscript Resolve

Miniscript, especificado no BIP 379, e uma representacao estruturada do Bitcoin Script. Ele mapeia um conjunto definido de fragmentos componives para opcodes de Script. Cada fragmento tem propriedades conhecidas: seu tipo, custo de recursos e dados de witness necessarios.

Como o mapeamento e estruturado, software pode analisar uma expressao Miniscript e determinar cada caminho de gasto, cada chave necessaria, cada restricao de timelock e o tamanho exato da witness para cada caminho. Duas expressoes que codificam a mesma politica podem ser comparadas e demonstradas como equivalentes.

Composicao se torna segura. Uma carteira pode pegar dois fragmentos Miniscript, combina-los com AND ou OR e calcular o consumo de recursos do script resultante antes de transmitir. Se qualquer caminho de execucao exceder os limites de consenso, o compilador rejeita a composicao no momento da construcao, nao no momento do gasto.
