## A Cadeia de Hashes

Cada cabecalho de bloco contem o hash do cabecalho do bloco anterior. O bloco 500.000 contem o hash do bloco 499.999, que contem o hash do bloco 499.998, e assim por diante ate o bloco genesis (bloco 0).

Mudar qualquer dado no bloco 499.999 mudaria seu hash. O bloco 500.000 entao conteria um hash de bloco anterior incorreto e se tornaria invalido. Para alterar um bloco historico, um atacante deve refazer o proof of work daquele bloco e de todos os blocos seguintes.

Essa conexao cumulativa e o que torna o blockchain a prova de adulteracao. Quanto mais profundo um bloco esta, mais trabalho e necessario para altera-lo. Cada novo bloco adiciona outra camada de protecao a cada bloco abaixo dele.
