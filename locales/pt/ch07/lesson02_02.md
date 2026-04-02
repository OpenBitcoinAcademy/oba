## Scripts Multi-assinatura

Um script multi-assinatura (multisig) exige M assinaturas de N chaves publicas para autorizar o gasto. Um multisig 2-de-3 exige quaisquer 2 de 3 chaves designadas.

Em um output multisig puro, o script de bloqueio contem: M <pubkey1> <pubkey2> ... <pubkeyN> N OP_CHECKMULTISIG. O script de desbloqueio fornece: OP_0 <sig1> <sig2>. O OP_0 inicial e uma solucao para um bug historico de off-by-one no OP_CHECKMULTISIG.

Na pratica, multisig e encapsulado em P2SH. O remetente ve um endereco P2SH padrao. Os detalhes do multisig ficam ocultos dentro do redeem script e sao revelados apenas no momento do gasto. Isso mantem os outputs compactos e a complexidade privada ate o momento de gastar.
