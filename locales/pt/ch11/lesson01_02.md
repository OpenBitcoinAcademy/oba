## O Mempool

Quando voce transmite uma transacao, ela nao entra em um bloco imediatamente. Ela entra no mempool: uma area de espera onde transacoes nao confirmadas ficam ate que um minerador as inclua em um bloco.

Cada no mantem seu proprio mempool. Transacoes se propagam pela rede em segundos, mas a confirmacao (inclusao em um bloco) leva em media 10 minutos. Em periodos movimentados, o mempool cresce e transacoes com taxas baixas podem esperar mais.

Uma transacao no mempool esta nao confirmada. Ela foi validada pelos nos (assinaturas corretas, inputs nao gastos), mas ainda nao foi escrita no blockchain.
