## Assinatura em Duas Rodadas

A assinatura FROST acontece em duas rodadas entre os assinantes participantes. Um coordenador retransmite mensagens, mas nunca aprende o suficiente para forjar uma assinatura ou extrair shares secretos. O coordenador nao e confiavel.

Na primeira rodada, cada assinante gera um par de nonces aleatorios e envia os commitments de nonce correspondentes ao coordenador. O coordenador coleta commitments de todos os t assinantes participantes e os transmite ao grupo. Esses commitments vinculam cada assinante a seus nonces antes que qualquer pessoa veja os valores de outros assinantes, impedindo que um assinante desonesto escolha nonces adaptativamente para manipular o resultado.

Na segunda rodada, cada assinante calcula o nonce agregado a partir dos commitments coletados, constroi o hash de desafio Schnorr e produz um share de assinatura usando seu share de chave secreta e seu nonce. Cada assinante envia seu share de assinatura ao coordenador.
