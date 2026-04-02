## Bloom Filters e seu Problema de Privacidade

O **BIP 37** introduziu Bloom filters como uma forma de clientes SPV solicitarem apenas transacoes relevantes de nos completos. Um Bloom filter e uma estrutura de dados probabilistica. O cliente cria um filtro contendo seus enderecos e o envia a um par. O par envia apenas transacoes que correspondem ao filtro.

O design tem uma falha seria de privacidade. O no completo ve o Bloom filter e pode deduzir quais enderecos pertencem ao cliente. Embora um Bloom filter tenha falsos positivos, pesquisas demonstraram que um no malicioso pode identificar os enderecos do cliente com alta precisao analisando o padrao de bits do filtro.

Um no servindo requisicoes de Bloom filter tambem suporta um custo computacional pesado. Ele deve testar cada transacao em cada bloco contra o filtro. Isso cria um vetor de negacao de servico: um atacante pode se conectar muitas vezes com filtros diferentes e forcar o no a fazer trabalho caro. Muitos operadores de nos agora desabilitam o suporte ao BIP 37.
