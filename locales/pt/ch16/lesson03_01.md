## Os Tres Mapas

Um PSBT contem tres tipos de mapas chave-valor: um mapa global, um mapa por input e um mapa por output. Cada mapa armazena pares chave-valor tipados terminados por um byte 0x00. O tipo de chave determina o que a entrada significa. O payload de chave e o payload de valor carregam os dados.

O mapa global contem dados que se aplicam a transacao inteira. No PSBTv0, a entrada global mais importante e UNSIGNED_TX (tipo de chave 0x00), que contem a transacao nao assinada completa. No PSBTv2, o mapa global contem TX_VERSION, FALLBACK_LOCKTIME, INPUT_COUNT e OUTPUT_COUNT como campos separados. A transacao nao assinada e reconstruida a partir de mapas por input e por output em vez de armazenada inteira.

Entradas globais tambem incluem XPUB (tipo de chave 0x01), que mapeia uma chave publica estendida para seu caminho de derivacao BIP 32. Isso permite que um assinante verifique que outputs de troco derivam da mesma raiz de carteira sem precisar de acesso ao descriptor completo da carteira.
