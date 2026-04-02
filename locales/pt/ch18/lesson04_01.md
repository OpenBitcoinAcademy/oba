## Abrindo um Canal

Abrir um canal requer uma transacao on-chain: a transacao de funding. Ela cria um output multisig 2-de-2. Antes de transmitir, ambas as partes assinam a primeira transacao de commitment (a rede de seguranca que devolve fundos se o canal falhar).

O canal esta aberto uma vez que a transacao de funding e confirmada. Ambas as partes podem comecar a atualizar o saldo trocando novas transacoes de commitment. O custo on-chain e uma taxa de transacao para a transacao de funding.
