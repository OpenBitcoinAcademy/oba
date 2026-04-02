## Alvo, Dificuldade e o Nonce

O alvo e um numero de 256 bits. Um hash de bloco valido deve ser numericamente menor que o alvo. Um alvo mais baixo significa que existem menos hashes validos, o que torna encontrar um mais dificil. A dificuldade e o inverso do alvo: quando o alvo desce, a dificuldade sobe.

O cabecalho do bloco contem um campo nonce de 32 bits. Mineradores incrementam esse valor a cada tentativa de hash, percorrendo todas as $2^{32}$ (cerca de 4,3 bilhoes) de possibilidades. Mineradores modernos esgotam o espaco do nonce em menos de um segundo.

Quando o espaco do nonce acaba, mineradores modificam outros campos para criar novas entradas de hash. A tecnica mais comum altera o campo extra nonce da transacao coinbase. Isso muda a raiz merkle no cabecalho do bloco, dando ao minerador um novo conjunto de $2^{32}$ nonces para tentar.
