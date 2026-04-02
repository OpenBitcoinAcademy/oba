## Byte por Byte

**Version** (4 bytes): 01000000 significa versao 1. 02000000 significa versao 2 (habilita restricoes de sequencia BIP68). Codificacao little-endian: o byte menos significativo vem primeiro.

**Inputs**: uma contagem varint seguida por cada input. Cada input contem o hash da transacao anterior (32 bytes, invertido), o indice do output (4 bytes), o comprimento do script de input (varint), o script de input (variavel) e o numero de sequencia (4 bytes).

**Outputs**: uma contagem varint seguida por cada output. Cada output contem o valor em satoshis (8 bytes, little-endian) e o comprimento do script de output (varint) seguido pelo script de output.

**Witness** (somente segwit): para cada input, uma contagem de itens de witness seguida pelo comprimento e dados de cada item. Inputs legacy tem zero itens de witness.

**Locktime** (4 bytes): normalmente 00000000. Um valor diferente de zero restringe quando a transacao pode ser minerada.

O ID da transacao (txid) e o hash double-SHA-256 da transacao serializada, com os dados de witness excluidos (para segwit) ou incluidos (para legacy).
