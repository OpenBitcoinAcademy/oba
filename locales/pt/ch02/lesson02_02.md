## Propagacao pela Rede

A carteira de Alice assina a transacao com sua chave privada e a transmite para a rede Bitcoin. Sua carteira se conecta a varios nos, que verificam a transacao de forma independente: formato correto, assinatura valida, inputs nao gastos, valores de output que nao excedem os valores de input.

Cada no que aceita a transacao a adiciona ao seu mempool (uma lista local de transacoes nao confirmadas) e a repassa para seus pares. Em segundos, a transacao alcanca milhares de nos pela rede. Nenhum servidor central coordena isso. Cada no segue as mesmas regras de forma independente.

A carteira de Bob monitora a rede em busca de transacoes pagando para seu endereco. Quando ela detecta a transacao de Alice, aparece como "nao confirmada" na carteira de Bob. O pagamento foi transmitido, mas ainda nao foi registrado em um bloco.
