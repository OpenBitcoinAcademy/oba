## Mineracao como Loteria Descentralizada

O Bitcoin nao tem uma autoridade central que decide quais transacoes entram nos blocos. Em vez disso, mineradores competem em uma loteria computacional. O vencedor ganha o direito de adicionar o proximo bloco a cadeia.

Cada minerador pega um conjunto de transacoes nao confirmadas, monta-as em um bloco candidato e aplica hash no cabecalho do bloco usando SHA-256 (aplicado duas vezes). O resultado deve ser um numero abaixo do alvo atual. A maioria das tentativas falha. O minerador muda o campo nonce no cabecalho e aplica hash novamente.

Esse processo se repete bilhoes de vezes por segundo em todos os mineradores do mundo. O minerador que encontra um hash valido primeiro transmite o bloco. Outros nos o verificam de forma independente. Ninguem precisa de permissao para minerar, e ninguem pode prever quem vencera.
