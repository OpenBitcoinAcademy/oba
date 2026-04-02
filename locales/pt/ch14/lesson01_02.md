## Por que o Tweak Funciona

O tweak vincula a arvore de scripts a chave publica criptograficamente. Mudar qualquer script na arvore muda a raiz Merkle, que muda o tweak, que muda Q. Um script comprometido em Q no momento da criacao do output nao pode ser modificado depois.

A chave privada e ajustada de forma identica: tweaked_privkey = (privkey + t) mod n. O detentor da chave privada interna pode calcular a chave privada ajustada e assinar diretamente. Esse e o gasto por key path.

Tagged hashes previnem colisoes entre diferentes usos da funcao hash. A tag "TapTweak" e separada por dominio: SHA256(SHA256("TapTweak") || SHA256("TapTweak") || data). Tags diferentes produzem saidas de hash diferentes para os mesmos dados, eliminando ataques entre protocolos.
