## A Loteria da Mineracao

Mineracao e uma loteria descentralizada. Cada minerador cria um bloco candidato a partir das transacoes em seu mempool. Eles aplicam hash no cabecalho do bloco repetidamente, mudando um numero chamado nonce a cada vez, procurando um valor de hash abaixo do alvo definido pela rede.

Encontrar um hash valido exige trilhoes de tentativas. Verificar um hash valido exige computa-lo uma vez. Essa assimetria e o nucleo do proof of work: criar um bloco valido e caro, verificar um e barato.

O minerador que encontra um hash valido primeiro transmite o bloco. Outros mineradores verificam, aceitam e comecam a trabalhar no proximo bloco. O vencedor recebe a recompensa do bloco: bitcoins recem-criados (o subsidio) mais a soma de todas as taxas de transacao no bloco.
