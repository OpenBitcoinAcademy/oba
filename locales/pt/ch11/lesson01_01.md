## O que e um Bloco?

Um bloco e um lote de transacoes agrupadas e adicionadas ao blockchain. Cada bloco contem um cabecalho e uma lista de transacoes.

O cabecalho do bloco tem seis campos: versao, hash do bloco anterior, raiz merkle (um hash de todas as transacoes no bloco), timestamp, alvo de dificuldade e nonce. O hash do bloco anterior conecta cada bloco ao anterior, formando uma cadeia.

Mineradores competem para encontrar um valor de nonce que faca o hash do cabecalho do bloco cair abaixo do alvo de dificuldade. Isso exige trilhoes de tentativas. O primeiro minerador a encontrar um nonce valido transmite o bloco. Outros nos o verificam e o adicionam a sua copia da cadeia.
