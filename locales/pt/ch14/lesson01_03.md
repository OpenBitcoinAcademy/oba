## Chaves Publicas X-Only

O Taproot usa chaves publicas X-only de 32 bytes em vez do formato comprimido tradicional de 33 bytes. A coordenada Y e escolhida implicitamente como par. Isso economiza um byte por chave na cadeia e em assinaturas.

Uma chave publica comprimida padrao tem um byte de prefixo (0x02 para Y par, 0x03 para Y impar) seguido de 32 bytes de coordenada X. O Taproot descarta o prefixo e exige que a coordenada Y seja par. Se o Q calculado tiver Y impar, a implementacao nega a chave privada (o que inverte Y para par) antes de assinar.

Essa convencao simplifica a verificacao em lote e a agregacao de chaves. Cada chave de output Taproot tem 32 bytes. Cada assinatura Schnorr tem 64 bytes. Nao ha codificacoes de comprimento variavel.
