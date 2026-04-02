## Varints e Endianness

O Bitcoin usa duas convencoes de codificacao em toda sua serializacao.

**Varints** (inteiros de comprimento variavel) codificam contagens e comprimentos. Um valor abaixo de 253 cabe em um byte. Valores de 253 a 65535 usam um prefixo 0xFD seguido de 2 bytes. Valores ate ~4 bilhoes usam 0xFE seguido de 4 bytes. Valores maiores usam 0xFF seguido de 8 bytes.

**Little-endian** coloca o byte menos significativo primeiro. A versao 1 e serializada como 01 00 00 00, nao 00 00 00 01. Valores em satoshis, alturas de bloco e a maioria dos campos numericos usam little-endian.

Hashes de transacao (txids) sao exibidos em ordem de bytes invertida por convencao. O hash double-SHA-256 produz bytes em uma ordem, mas exploradores de blocos e o Bitcoin Core os exibem invertidos. Isso e uma convencao de exibicao, nao uma regra de serializacao.
