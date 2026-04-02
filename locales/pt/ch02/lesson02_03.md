## Do Mempool ao Blockchain

A transacao de Alice fica no mempool, esperando que um minerador a inclua em um bloco. Mineradores selecionam transacoes de seu mempool, priorizando aquelas com maiores taxas por unidade de peso.

Cerca de cinco minutos depois, um minerador encontra uma solucao valida de proof of work para um novo bloco que inclui a transacao de Alice. O minerador transmite o bloco. Cada no completo verifica o bloco de forma independente: proof of work valido, transacoes validas, formato correto. Os nos que aceitam o bloco o adicionam a sua copia do blockchain.

A transacao de Alice agora tem uma confirmacao. A carteira de Bob atualiza o status. Cada bloco subsequente construido sobre este adiciona outra confirmacao. Apos seis confirmacoes (cerca de uma hora), a transacao e considerada altamente segura contra reversao.
