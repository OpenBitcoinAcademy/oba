## De Sementes a Arvores

Uma semente e um numero aleatorio, tipicamente de 128 a 256 bits. E a raiz de toda derivacao de chaves. A partir dessa semente, uma hierarquia de chaves pode ser gerada: a semente produz uma chave-mestra, a chave-mestra produz chaves filhas, cada filha pode produzir suas proprias filhas.

Essa estrutura em arvore e a base das carteiras HD (Hierarchical Deterministic), definidas no BIP32. A arvore permite organizar chaves por proposito: um ramo para receber pagamentos, outro para troco, outro para uma conta diferente.

A semente nunca muda. Cada chave na arvore pode ser regenerada a partir dela. Um unico backup protege um numero ilimitado de chaves futuras.
