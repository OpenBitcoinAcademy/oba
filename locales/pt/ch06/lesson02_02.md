## Construindo a Transacao

Uma transacao legacy tem quatro campos: version, inputs, outputs e locktime. Transacoes segwit modernas adicionam tres mais: um marker, um flag e uma estrutura de witness que guarda os dados de autorizacao (assinaturas) separadamente dos inputs.

**Version** e um numero (atualmente 1 ou 2) que diz aos nos quais regras de validacao aplicar.

**Inputs** listam os UTXOs sendo gastos. Cada input especifica o ID da transacao anterior, o indice do output naquela transacao, um script de input e um numero de sequencia.

**Outputs** listam os novos UTXOs sendo criados. Cada output especifica um valor em satoshis e um script de bloqueio.

**Locktime** normalmente e zero. Quando definido para uma altura de bloco ou timestamp futuros, a transacao nao pode ser incluida em um bloco ate aquele ponto.

A transacao e serializada em uma sequencia de bytes, recebe hash duplo com SHA-256, e o resultado e o ID da transacao.
