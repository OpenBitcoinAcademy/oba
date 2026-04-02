## Inputs Referenciam Outputs Anteriores

Cada input aponta para um output especifico de uma transacao anterior. Ele faz isso com duas informacoes: o ID da transacao (um hash) e o indice do output (qual output naquela transacao).

Quando voce gasta bitcoin, voce prova que controla a chave que pode desbloquear um output anterior. Em transacoes legacy, a prova (assinatura e chave publica) fica no script de input. Em transacoes segwit modernas, a prova fica em uma estrutura de witness separada, e o script de input esta vazio.

Uma vez que um output e referenciado por um input, ele esta gasto. Nao pode ser gasto novamente. E assim que o Bitcoin previne gasto duplo sem uma autoridade central.
