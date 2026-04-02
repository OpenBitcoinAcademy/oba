## Como uma Transacao se Parece

Uma transacao Bitcoin e uma estrutura de dados que transfere bitcoin de um dono para outro. Ela contem inputs (que referenciam outputs anteriores a serem gastos), outputs (que criam novos valores gastaveis) e dados de autorizacao (assinaturas provando que o gastador controla as chaves).

Inputs dizem de onde o bitcoin vem. Cada input referencia um output de uma transacao anterior que ainda nao foi gasto. Esse output nao gasto e chamado UTXO (unspent transaction output).

Outputs dizem para onde o bitcoin vai. Cada output especifica um valor em satoshis e uma condicao de bloqueio (um script) que determina quem pode gasta-lo.

Uma transacao consome outputs antigos e cria novos. Nada e "armazenado em uma conta". O Bitcoin rastreia propriedade por uma cadeia de outputs, cada um vinculado a uma chave especifica.
