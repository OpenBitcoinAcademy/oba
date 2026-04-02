## Provas Merkle

Uma prova Merkle permite verificar que uma transacao especifica esta incluida em um bloco sem baixar cada transacao. A prova consiste no hash da transacao, nos hashes irmaos ao longo do caminho ate a raiz e na raiz Merkle do cabecalho do bloco.

Para verificar, comece com o hash da transacao. Faca hash dele com seu irmao. Faca hash do resultado com o proximo irmao. Continue ate chegar a raiz. Se sua raiz calculada corresponder a raiz Merkle no cabecalho do bloco, a transacao esta no bloco.

Um bloco com 4.000 transacoes requer uma prova Merkle de apenas cerca de 12 hashes (log2(4000)). Clientes SPV usam provas Merkle para verificar transacoes sem baixar blocos completos.
